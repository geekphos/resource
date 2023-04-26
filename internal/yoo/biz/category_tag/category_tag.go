package category_tag

import (
	"context"

	"phos.cc/yoo/internal/yoo/store"
	v1 "phos.cc/yoo/pkg/api/yoo/v1"
)

type CategoryTagBiz interface {
	Tree(ctx context.Context) ([]*v1.AllCategoryTagResponse, error)
	Categories(ctx context.Context) ([]string, error)
	Tags(ctx context.Context) ([]string, error)
}

type categoryTagBiz struct {
	ds store.IStore
}

var _ CategoryTagBiz = (*categoryTagBiz)(nil)

func New(ds store.IStore) *categoryTagBiz {
	return &categoryTagBiz{ds: ds}
}

func (b *categoryTagBiz) Tree(ctx context.Context) ([]*v1.AllCategoryTagResponse, error) {
	return b.ds.CategoryTags().Tree(ctx)
}

func (b *categoryTagBiz) Categories(ctx context.Context) ([]string, error) {
	return b.ds.CategoryTags().Categories(ctx)
}

func (b *categoryTagBiz) Tags(ctx context.Context) ([]string, error) {
	return b.ds.CategoryTags().Tags(ctx)
}
