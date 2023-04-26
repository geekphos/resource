package resource

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
	"gorm.io/datatypes"
	"sort"

	"phos.cc/yoo/internal/pkg/errno"
	"phos.cc/yoo/internal/pkg/model"
	"phos.cc/yoo/internal/yoo/store"
	v1 "phos.cc/yoo/pkg/api/yoo/v1"
)

type ResourceBiz interface {
	Update(ctx context.Context, r *v1.UpdateResourceRequest) error
	List(ctx context.Context, r *v1.ListResourceRequest) ([]*v1.ListResourceResponse, int64, error)
	Get(ctx context.Context, id int32) (*v1.GetResourceResponse, error)
	Used(ctx context.Context, r *v1.ListUsedResourceRequest) ([]map[string][]*v1.ListUsedResourceResponse, []string, error)
}

type resourceBiz struct {
	ds store.IStore
}

var _ ResourceBiz = (*resourceBiz)(nil)

func New(ds store.IStore) *resourceBiz {
	return &resourceBiz{ds: ds}
}

func (b *resourceBiz) Update(ctx context.Context, r *v1.UpdateResourceRequest) error {
	resourceM, err := b.ds.Resources().Get(ctx, r.ID)
	if err != nil {
		return errno.InternalServerError
	}
	_ = copier.Copy(resourceM, r)

	if err := b.ds.Resources().Update(ctx, resourceM); err != nil {
		return errno.InternalServerError
	}
	return nil
}

func (b *resourceBiz) List(ctx context.Context, r *v1.ListResourceRequest) ([]*v1.ListResourceResponse, int64, error) {
	var resourceM = &model.ResourceM{}
	_ = copier.Copy(resourceM, r)

	if r.Tag != "" {
		resourceM.Tags = datatypes.JSON(`["` + r.Tag + `"]`)
	}

	resources, total, err := b.ds.Resources().List(ctx, r.Page, r.PageSize, resourceM)
	if err != nil {
		return nil, 0, errno.InternalServerError
	}

	var resourcesR []*v1.ListResourceResponse
	for _, resource := range resources {
		var resourceR = &v1.ListResourceResponse{}
		_ = copier.Copy(resourceR, resource)
		resourcesR = append(resourcesR, resourceR)
	}

	return resourcesR, total, nil
}

func (b *resourceBiz) Get(ctx context.Context, id int32) (*v1.GetResourceResponse, error) {
	resource, err := b.ds.Resources().Get(ctx, id)
	if err != nil {
		return nil, errno.ErrResourceNotFound
	}

	var resourceR = &v1.GetResourceResponse{}
	_ = copier.Copy(resourceR, resource)

	return resourceR, nil
}

func (b *resourceBiz) Used(ctx context.Context, r *v1.ListUsedResourceRequest) ([]map[string][]*v1.ListUsedResourceResponse, []string, error) {
	menus, err := b.ds.Menus().All(ctx, &model.MenuM{MenuType: 2})
	if err != nil {
		return nil, nil, errno.InternalServerError
	}

	var ids []int32
	var letters []string

	// 过滤掉非资源菜单
	menus = lo.Filter(menus, func(menu *model.MenuM, _ int) bool {
		return menu.ResourceID != nil
	})

	for _, menu := range menus {
		letters = append(letters, menu.Letter)
	}

	menus = lo.Filter(menus, func(menu *model.MenuM, _ int) bool {
		if r.Letter != "" {
			return menu.Letter == r.Letter
		}
		return true
	})

	for _, menu := range menus {
		ids = append(ids, *menu.ResourceID)
	}

	resources, err := b.ds.Resources().GetUsedResource(ctx, ids, r.Name, r.Tags)
	if err != nil {
		return nil, nil, errno.InternalServerError
	}

	var rm = make(map[int32]*model.ResourceM)

	for _, resource := range resources {
		rm[resource.ID] = resource
	}

	var res = make(map[string][]*v1.ListUsedResourceResponse)

	for _, menu := range menus {
		if resource, ok := rm[*menu.ResourceID]; ok {
			res[menu.Letter] = append(res[menu.Letter], &v1.ListUsedResourceResponse{
				ID:          resource.ID,
				Name:        menu.Name,
				Description: resource.Label,
				Badge:       menu.Icon,
				CreatedAt:   resource.CreatedAt,
				UpdatedAt:   resource.UpdatedAt,
			})
		}
	}

	// convert res to array and sort by letter
	var resArr []map[string][]*v1.ListUsedResourceResponse

	for letter, resources := range res {
		resArr = append(resArr, map[string][]*v1.ListUsedResourceResponse{
			letter: resources,
		})
	}

	// 对字母进行排序
	sort.Strings(letters)

	// use sort.Slice to sort
	sort.Slice(resArr, func(i, j int) bool {
		return lo.Keys(resArr[i])[0] < lo.Keys(resArr[j])[0]
	})

	return resArr, letters, nil
}
