package store

import (
	"context"

	"gorm.io/gorm"

	"phos.cc/yoo/internal/pkg/model"
)

type MenuStore interface {
	Create(ctx context.Context, menu *model.MenuM) error
	Update(ctx context.Context, menu *model.MenuM) error
	All(ctx context.Context, menu *model.MenuM) ([]*model.MenuM, error)
	Get(ctx context.Context, id int32) (*model.MenuM, error)
	Delete(ctx context.Context, id int32) error
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
