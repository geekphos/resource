package store

import (
	"context"
	"encoding/json"
	"gorm.io/datatypes"

	"gorm.io/gorm"

	"phos.cc/yoo/internal/pkg/model"
)

type MenuStore interface {
	Create(ctx context.Context, menu *model.MenuM) error
	Update(ctx context.Context, menu *model.MenuM) error
	All(ctx context.Context, menu *model.MenuM) ([]*model.MenuM, error)
	Get(ctx context.Context, id int32) (*model.MenuM, error)
	GetByName(ctx context.Context, name string) (*model.MenuM, error)
	Delete(ctx context.Context, id int32) error
	GetLeaveMenus(ctx context.Context) ([]*model.MenuM, error)
	GetLeaveMenusWithCond(ctx context.Context, menu *model.MenuM) ([]*model.MenuM, error)
	GetMenuByIds(ctx context.Context, ids []int32) ([]*model.MenuM, error)
}

type menuStore struct {
	db *gorm.DB
}

var _ MenuStore = (*menuStore)(nil)

func newMenus(db *gorm.DB) MenuStore {
	return &menuStore{db: db}
}

func (m *menuStore) Create(ctx context.Context, menu *model.MenuM) error {
	return m.db.WithContext(ctx).Create(menu).Error
}

func (m *menuStore) Update(ctx context.Context, menu *model.MenuM) error {
	return m.db.WithContext(ctx).Model(menu).Save(menu).Error
}

func (m *menuStore) All(ctx context.Context, menu *model.MenuM) ([]*model.MenuM, error) {
	var menus []*model.MenuM
	var name string
	if menu.Name != "" {
		name = "%" + menu.Name + "%"
		menu.Name = ""
	}
	query := m.db.WithContext(ctx).Model(menu)
	if name != "" {
		query = query.Where("name like ?", name)
	}
	if menu.Letter != "" {
		query = query.Where("letter = ?", menu.Letter)
	}
	err := query.Find(&menus).Error
	return menus, err
}

func (m *menuStore) Get(ctx context.Context, id int32) (*model.MenuM, error) {
	var menu model.MenuM
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&menu).Error
	return &menu, err
}

func (m *menuStore) Delete(ctx context.Context, id int32) error {
	return m.db.WithContext(ctx).Where("id = ?", id).Delete(&model.MenuM{}).Error
}

func (m *menuStore) GetByName(ctx context.Context, name string) (*model.MenuM, error) {
	var menu model.MenuM
	err := m.db.WithContext(ctx).Where("name = ?", name).First(&menu).Error
	return &menu, err
}

func (m *menuStore) GetLeaveMenus(ctx context.Context) ([]*model.MenuM, error) {
	var menus []*model.MenuM
	err := m.db.WithContext(ctx).Where("id not in (select parent_id from menus where parent_id is not null)").Find(&menus).Error
	return menus, err
}

func (m *menuStore) GetLeaveMenusWithCond(ctx context.Context, menu *model.MenuM) ([]*model.MenuM, error) {
	var menus []*model.MenuM
	query := m.db.WithContext(ctx).Where("id not in (select parent_id from menus where parent_id is not null)")
	if menu.Letter != "" {
		query = query.Where("letter = ?", menu.Letter)
	}

	if menu.Name != "" {
		query = query.Where("name like ?", "%"+menu.Name+"%")
	}

	if menu.Categories != nil {

		// convert menu.Categories to []byte

		// convert categories to []int
		var categories []int32
		_ = json.Unmarshal(menu.Categories, &categories)
		for _, category := range categories {
			query = query.Where(datatypes.JSONArrayQuery("categories").Contains(category))
		}
	}

	err := query.Find(&menus).Error
	return menus, err
}

func (m *menuStore) GetMenuByIds(ctx context.Context, ids []int32) ([]*model.MenuM, error) {
	var menus []*model.MenuM
	err := m.db.WithContext(ctx).Where("id in ?", ids).Find(&menus).Error
	return menus, err
}
