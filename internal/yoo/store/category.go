package store

import (
	"context"

	"gorm.io/gorm"
	"phos.cc/yoo/internal/pkg/model"
)

type CategoryStore interface {
	Create(ctx context.Context, category *model.CategoryM) error
	Update(ctx context.Context, category *model.CategoryM) error
	Get(ctx context.Context, id int32) (*model.CategoryM, error)
	All(ctx context.Context) ([]*model.CategoryM, error)
	Delete(ctx context.Context, id int32) error
}

type categoryStore struct {
	db *gorm.DB
}

var _ CategoryStore = (*categoryStore)(nil)

func newCategories(db *gorm.DB) CategoryStore {
	return &categoryStore{db: db}
}

func (c *categoryStore) Create(ctx context.Context, category *model.CategoryM) error {
	return c.db.WithContext(ctx).Create(category).Error
}

func (c *categoryStore) Update(ctx context.Context, category *model.CategoryM) error {
	return c.db.WithContext(ctx).Model(category).Save(category).Error
}

func (c *categoryStore) All(ctx context.Context) ([]*model.CategoryM, error) {
	var categories []*model.CategoryM
	err := c.db.WithContext(ctx).Find(&categories).Error
	return categories, err
}

func (c *categoryStore) Delete(ctx context.Context, id int32) error {
	return c.db.WithContext(ctx).Delete(&model.CategoryM{}, id).Error
}

func (c *categoryStore) Get(ctx context.Context, id int32) (*model.CategoryM, error) {
	var category model.CategoryM
	err := c.db.WithContext(ctx).Where("id = ?", id).First(&category).Error
	return &category, err
}
