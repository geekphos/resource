package menu

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/samber/lo"
	"gorm.io/datatypes"

	"github.com/jinzhu/copier"
	"github.com/mozillazg/go-pinyin"

	"phos.cc/yoo/internal/pkg/errno"
	"phos.cc/yoo/internal/pkg/model"
	"phos.cc/yoo/internal/yoo/store"
	v1 "phos.cc/yoo/pkg/api/yoo/v1"
)

type MenuBiz interface {
	Create(ctx context.Context, r *v1.CreateMenuRequest) error
	Update(ctx context.Context, r *v1.UpdateMenuRequest) error
	Updates(ctx context.Context, rl []*v1.UpdateMenuRequest) error
	Get(ctx context.Context, id int32) (*v1.GetMenuResponse, error)
	Tree(ctx context.Context, r *v1.ListMenuRequest) ([]*v1.ListMenuResponse, error)
	Delete(ctx context.Context, id int32) error
	GetMenuPath(ctx context.Context, name string) (string, error)
	GetLeaveMenus(ctx context.Context, r *v1.GetLeaveMenuRequest) ([]map[string][]*v1.GetLeaveMenuResponse, []string, error)
	GetMenuByIds(ctx context.Context, ids []int32) ([]*v1.GetMenuResponse, error)
}

type menuBiz struct {
	ds store.IStore
}

var _ MenuBiz = (*menuBiz)(nil)
var pyArgs = pinyin.NewArgs()

func New(ds store.IStore) MenuBiz {
	return &menuBiz{ds: ds}
}

func (b *menuBiz) Create(ctx context.Context, r *v1.CreateMenuRequest) error {
	var menuM = &model.MenuM{}
	_ = copier.Copy(menuM, r)

	menuM.Letter = getAlphaLetter(r.Name)

	if err := b.ds.Menus().Create(ctx, menuM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key '(plan_id|project_id)'", err.Error()); match {
			return errno.ErrMenuAlreadyExist
		} else {
			return errno.InternalServerError
		}
	}

	return nil
}

func (b *menuBiz) Update(ctx context.Context, r *v1.UpdateMenuRequest) error {
	menuM, err := b.ds.Menus().Get(ctx, r.ID)
	if err != nil {
		return errno.ErrMenuNotFound
	}
	_ = copier.Copy(menuM, r)

	if r.Name != nil {
		menuM.Letter = getAlphaLetter(*r.Name)
	}

	if r.ParentID != nil && *r.ParentID == 0 {
		menuM.ParentID = nil
	}

	if r.ResourceID != nil && *r.ResourceID == 0 {
		menuM.ResourceID = nil
	}

	if err := b.ds.Menus().Update(ctx, menuM); err != nil {
		return errno.InternalServerError
	}
	return nil
}

func (b *menuBiz) Updates(ctx context.Context, rl []*v1.UpdateMenuRequest) error {

	tds := b.ds.TX()

	for _, r := range rl {
		menuM, err := tds.Menus().Get(ctx, r.ID)
		if err != nil {
			return errno.ErrMenuNotFound
		}
		_ = copier.Copy(menuM, r)

		if r.ParentID != nil && *r.ParentID == 0 {
			menuM.ParentID = nil
		}

		if r.ResourceID != nil && *r.ResourceID == 0 {
			menuM.ResourceID = nil
		}
		if err := tds.Menus().Update(ctx, menuM); err != nil {
			tds.Rollback()
			return errno.InternalServerError
		}
	}

	if err := tds.Commit(); err != nil {
		return errno.InternalServerError
	}

	return nil
}

func (b *menuBiz) Get(ctx context.Context, id int32) (*v1.GetMenuResponse, error) {
	menuM, err := b.ds.Menus().Get(ctx, id)
	if err != nil {
		return nil, errno.ErrMenuNotFound
	}

	var resp = &v1.GetMenuResponse{}
	_ = copier.Copy(resp, menuM)

	return resp, nil
}

func (b *menuBiz) Tree(ctx context.Context, r *v1.ListMenuRequest) ([]*v1.ListMenuResponse, error) {
	var menuM = &model.MenuM{}
	_ = copier.Copy(menuM, r)

	ms, err := b.ds.Menus().All(ctx, menuM)
	if err != nil {
		return nil, errno.InternalServerError
	}
	return buildMenuRespTree(ms), nil
}

func buildMenuRespTree(ms []*model.MenuM) []*v1.ListMenuResponse {
	var menuMap = make(map[int32]*v1.ListMenuResponse)
	var rootMenus []*v1.ListMenuResponse

	for _, m := range ms {
		menuResp := &v1.ListMenuResponse{}
		_ = copier.Copy(menuResp, m)
		menuMap[m.ID] = menuResp
	}

	for _, m := range ms {
		if m.ParentID == nil {
			menuMap[m.ID].Depth = 0
			rootMenus = append(rootMenus, menuMap[m.ID])
		} else {
			parent := menuMap[*m.ParentID]
			menuMap[m.ID].Depth = parent.Depth + 1
			parent.Children = append(parent.Children, menuMap[m.ID])
		}
	}

	return sortMenuRespTree(rootMenus)
}

// sortMenuRespTree sorts the menu tree by the number field.
func sortMenuRespTree(mt []*v1.ListMenuResponse) []*v1.ListMenuResponse {
	for _, m := range mt {
		if len(m.Children) > 0 {
			sortMenuRespTree(m.Children)
		}
	}

	sort.Slice(mt, func(i, j int) bool {
		return mt[i].Number < mt[j].Number
	})

	return mt

}

func (b *menuBiz) Delete(ctx context.Context, id int32) error {
	return b.ds.Menus().Delete(ctx, id)
}

func (b *menuBiz) GetMenuPath(ctx context.Context, name string) (string, error) {

	var menus []*model.MenuM

	menuM, err := b.ds.Menus().GetByName(ctx, name)
	if err != nil {
		return "", errno.ErrMenuNotFound
	}

	menus = append(menus, menuM)

	for menuM.ParentID != nil {
		menuM, err = b.ds.Menus().Get(ctx, *menuM.ParentID)
		if err != nil {
			return "", errno.ErrMenuNotFound
		}
		menus = append(menus, menuM)
	}

	var menuPath string
	for i := len(menus) - 1; i >= 0; i-- {
		menuPath += menus[i].Href
	}

	return menuPath, nil
}

func getAlphaLetter(name string) string {
	pyList := pinyin.Pinyin(name, pyArgs)

	if len(pyList) > 0 && len(pyList[0]) > 0 {
		alphaLetter := pyList[0][0]
		return strings.ToUpper(alphaLetter[:1])
	}

	return ""
}
func (b *menuBiz) GetLeaveMenus(ctx context.Context, r *v1.GetLeaveMenuRequest) ([]map[string][]*v1.GetLeaveMenuResponse, []string, error) {
	var menuM = &model.MenuM{}
	_ = copier.Copy(menuM, r)

	if r.Categories != nil {
		menuM.Categories = datatypes.JSON(fmt.Sprintf(`[%s]`, strings.Join(r.Categories, ",")))
	}

	ms, err := b.ds.Menus().GetLeaveMenus(ctx)
	if err != nil {
		return nil, nil, errno.InternalServerError
	}

	// letter map
	lm := make(map[string]struct{})

	for _, m := range ms {
		if m.Letter != "" {
			lm[m.Letter] = struct{}{}
		}
	}

	letters := lo.MapToSlice(lm, func(key string, _ struct{}) string {
		return key
	})

	// sort letters
	sort.Strings(letters)

	// get menu by condition
	ms, err = b.ds.Menus().GetLeaveMenusWithCond(ctx, menuM)

	if err != nil {
		return nil, nil, errno.InternalServerError
	}

	var menus []*v1.GetLeaveMenuResponse
	for _, m := range ms {
		menuResp := &v1.GetLeaveMenuResponse{}
		_ = copier.Copy(menuResp, m)
		menus = append(menus, menuResp)
	}

	var menuMap = make(map[string][]*v1.GetLeaveMenuResponse)

	for _, m := range menus {
		letter := m.Letter
		menuMap[letter] = append(menuMap[letter], m)
	}

	var resp []map[string][]*v1.GetLeaveMenuResponse

	for _, letter := range letters {
		if len(menuMap[letter]) != 0 {
			resp = append(resp, map[string][]*v1.GetLeaveMenuResponse{
				letter: menuMap[letter],
			})
		}
	}

	return resp, letters, nil
}

func (b *menuBiz) GetMenuByIds(ctx context.Context, ids []int32) ([]*v1.GetMenuResponse, error) {
	menuMs, err := b.ds.Menus().GetMenuByIds(ctx, ids)
	if err != nil {
		return nil, errno.InternalServerError
	}
	var resp []*v1.GetMenuResponse

	for _, m := range menuMs {
		menuResp := &v1.GetMenuResponse{}
		_ = copier.Copy(menuResp, m)
		resp = append(resp, menuResp)
	}

	return resp, nil
}
