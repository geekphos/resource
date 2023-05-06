package store

import (
	"context"
	"gorm.io/gorm"

	"phos.cc/yoo/internal/pkg/model"
)

type ResourceStore interface {
	Create(ctx context.Context, resource *model.ResourceM) error
	Update(ctx context.Context, resource *model.ResourceM) error
	List(ctx context.Context, page, pageSize int, resource *model.ResourceM) ([]*model.ResourceM, int64, error)
	All(ctx context.Context, m *model.ResourceM) ([]*model.ResourceM, error)
	Get(ctx context.Context, id int32) (*model.ResourceM, error)
}

type resourceStore struct {
	db *gorm.DB
}

var _ ResourceStore = (*resourceStore)(nil)

func newResources(db *gorm.DB) *resourceStore {
	return &resourceStore{db: db}
}

func (s *resourceStore) Create(ctx context.Context, resource *model.ResourceM) error {
	return s.db.WithContext(ctx).Create(resource).Error
}

func (s *resourceStore) Update(ctx context.Context, resource *model.ResourceM) error {
	return s.db.WithContext(ctx).Model(resource).Save(resource).Error
}

func (s *resourceStore) List(ctx context.Context, page, pageSize int, resource *model.ResourceM) ([]*model.ResourceM, int64, error) {
	var (
		resources []*model.ResourceM
		count     int64
	)

	query := s.db.WithContext(ctx).Model(&model.ResourceM{})
	if resource.Description != "" {
		query = query.Where("description LIKE ?", "%"+resource.Description+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&resources).Error; err != nil {
		return nil, 0, err
	}

	return resources, count, nil
}

func (s *resourceStore) Get(ctx context.Context, id int32) (*model.ResourceM, error) {
	var resource model.ResourceM
	err := s.db.WithContext(ctx).First(&resource, id).Error
	return &resource, err
}

func (s *resourceStore) All(ctx context.Context, m *model.ResourceM) ([]*model.ResourceM, error) {
	var resources []*model.ResourceM

	query := s.db.WithContext(ctx).Model(&model.ResourceM{})

	if m.Description != "" {
		query = query.Where("description LIKE ?", "%"+m.Description+"%")
	}

	err := query.Find(&resources).Error
	return resources, err
}
