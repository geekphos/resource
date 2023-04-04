package category_tag

import (
	"context"

	"phos.cc/yoo/internal/yoo/store"
	v1 "phos.cc/yoo/pkg/api/yoo/v1"
)

type CategoryTagBiz interface {
	All(ctx context.Context) ([]*v1.AllCategoryTagResponse, error)
}

type categoryTagBiz struct {
	ds store.IStore
}

var _ CategoryTagBiz = (*categoryTagBiz)(nil)

func New(ds store.IStore) *categoryTagBiz {
	return &categoryTagBiz{ds: ds}
}

func (b *categoryTagBiz) All(ctx context.Context) ([]*v1.AllCategoryTagResponse, error) {
	return b.ds.CategoryTags().All(ctx)
}
