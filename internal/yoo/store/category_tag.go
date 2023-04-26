package store

import (
	"context"
	"encoding/json"
	"gorm.io/datatypes"

	"github.com/samber/lo"
	"gorm.io/gorm"

	v1 "phos.cc/yoo/pkg/api/yoo/v1"
)

type CategoryTagStore interface {
	Tree(ctx context.Context) ([]*v1.AllCategoryTagResponse, error)
	Categories(ctx context.Context) ([]string, error)
	Tags(ctx context.Context) ([]string, error)
}

type categoryStore struct {
	db *gorm.DB
}

var _ CategoryTagStore = (*categoryStore)(nil)

func newCategoryTags(db *gorm.DB) *categoryStore {
	return &categoryStore{db: db}
}

func (c *categoryStore) Tree(ctx context.Context) ([]*v1.AllCategoryTagResponse, error) {
	var res []*v1.CategoryTag
	if result := c.db.Table("resources").Select([]string{"category", "tags"}).Scan(&res); result.Error != nil {
		return nil, result.Error
	}
	list := buildCategoryTree(res)
	return list, nil
}

func buildCategoryTree(category_tags []*v1.CategoryTag) []*v1.AllCategoryTagResponse {
	var res []*v1.AllCategoryTagResponse
	lo.ForEach(category_tags, func(ct *v1.CategoryTag, _ int) {
		if r, ok := lo.Find(res, func(r *v1.AllCategoryTagResponse) bool {
			return ct.Category == r.Name
		}); ok {
			var tags []string
			if err := json.Unmarshal(ct.Tags, &tags); err == nil {
				// r.Children = lo.Union(r.Children, tags)
				r.Children = lo.Union(r.Children, tags)
			}
		} else {
			var tags []string
			if err := json.Unmarshal(ct.Tags, &tags); err == nil {
				res = append(res, &v1.AllCategoryTagResponse{
					Name:     ct.Category,
					Children: tags,
				})
			}
		}
	})

	return res
}

func (c *categoryStore) Categories(ctx context.Context) ([]string, error) {
	var res []string
	if result := c.db.Table("resources").Select([]string{"category"}).Scan(&res); result.Error != nil {
		return nil, result.Error
	}
	return lo.Union(res), nil
}

func (c *categoryStore) Tags(ctx context.Context) ([]string, error) {
	var tags []datatypes.JSON
	if result := c.db.Table("resources").Select([]string{"tags"}).Scan(&tags); result.Error != nil {
		return nil, result.Error
	}
	var res [][]string
	lo.ForEach(tags, func(item datatypes.JSON, _ int) {
		var list []string
		if err := json.Unmarshal(item, &list); err == nil {
			res = append(res, list)
		}
	})
	return lo.Union(lo.Flatten(res)), nil
}
