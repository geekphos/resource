package category

import (
	"context"

	"github.com/jinzhu/copier"

	"phos.cc/yoo/internal/pkg/errno"
	"phos.cc/yoo/internal/pkg/model"
	"phos.cc/yoo/internal/yoo/store"
	v1 "phos.cc/yoo/pkg/api/yoo/v1"
)

type CategoryBiz interface {
	Create(ctx context.Context, r *v1.CreateCategoryRequest) error
	Update(ctx context.Context, r *v1.UpdateCategoryRequest) error
	All(ctx context.Context) ([]*v1.AllCategoryResponse, error)
}

type categoryBiz struct {
	ds store.IStore
}

var _ CategoryBiz = (*categoryBiz)(nil)

func New(ds store.IStore) CategoryBiz {
	return &categoryBiz{ds: ds}
}

func (b *categoryBiz) Create(ctx context.Context, r *v1.CreateCategoryRequest) error {
	var categoryM = &model.CategoryM{}
	_ = copier.Copy(categoryM, r)

	if err := b.ds.Categories().Create(ctx, categoryM); err != nil {
		return errno.InternalServerError
	}

	return nil
}

func (b *categoryBiz) Update(ctx context.Context, r *v1.UpdateCategoryRequest) error {
	categoryM, err := b.ds.Categories().Get(ctx, r.ID)
	if err != nil {
		return errno.ErrCategoryNotFound
	}

	_ = copier.Copy(categoryM, r)

	if err := b.ds.Categories().Update(ctx, categoryM); err != nil {
		return errno.InternalServerError
	}

	return nil
}

func (b *categoryBiz) All(ctx context.Context) ([]*v1.AllCategoryResponse, error) {
	categoryMs, err := b.ds.Categories().All(ctx)
	if err != nil {
		return nil, errno.InternalServerError
	}

	return buildCategoryTree(categoryMs), nil

}

func buildCategoryTree(categories []*model.CategoryM) []*v1.AllCategoryResponse {

	var categoryMap = make(map[int32]*v1.AllCategoryResponse)
	var rootCategories []*v1.AllCategoryResponse

	for _, c := range categories {
		categoryResp := &v1.AllCategoryResponse{}
		copier.Copy(categoryResp, c)
		categoryMap[c.ID] = categoryResp
	}

	for _, c := range categories {
		if c.ParentID == nil {
			rootCategories = append(rootCategories, categoryMap[c.ID])
		} else {
			parent := categoryMap[*c.ParentID]
			parent.Children = append(parent.Children, categoryMap[c.ID])
		}
	}

	return rootCategories

}
