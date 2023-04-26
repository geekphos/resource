package store

import (
	"context"
	"encoding/json"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"phos.cc/yoo/internal/pkg/model"
)

type ResourceStore interface {
	Update(ctx context.Context, resource *model.ResourceM) error
	List(ctx context.Context, page, pageSize int, resource *model.ResourceM) ([]*model.ResourceM, int64, error)
	Get(ctx context.Context, id int32) (*model.ResourceM, error)
	GetUsedResource(ctx context.Context, ids []int32, name string, tags []string) ([]*model.ResourceM, error)
}

type resourceStore struct {
	db *gorm.DB
}

var _ ResourceStore = (*resourceStore)(nil)

func newResources(db *gorm.DB) *resourceStore {
	return &resourceStore{db: db}
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
	if resource.Name != "" {
		query = query.Where("name LIKE ?", "%"+resource.Name+"%")
	}
	if resource.Label != "" {
		query = query.Where("label LIKE ?", "%"+resource.Label+"%")
	}

	if resource.Category != "" {
		query = query.Where("category = ?", resource.Category)
	}

	if resource.Tags != nil {
		var tags []string
		if err := json.Unmarshal(resource.Tags, &tags); err != nil {
			return nil, 0, err
		}
		for _, tag := range tags {
			query = query.Where(datatypes.JSONArrayQuery("tags").Contains(tag))
		}
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

func (s *resourceStore) GetUsedResource(ctx context.Context, ids []int32, name string, tags []string) ([]*model.ResourceM, error) {
	var resources []*model.ResourceM
	query := s.db.WithContext(ctx).Model(&model.ResourceM{})

	if len(ids) > 0 {
		query = query.Where("id in ?", ids)
	}

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if len(tags) > 0 {
		for _, tag := range tags {
			query = query.Where(datatypes.JSONArrayQuery("tags").Contains(tag))
		}
	}

	err := query.Find(&resources).Error
	return resources, err
}
