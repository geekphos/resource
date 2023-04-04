package biz

import (
	"phos.cc/yoo/internal/yoo/biz/category_tag"
	"phos.cc/yoo/internal/yoo/biz/menu"
	"phos.cc/yoo/internal/yoo/biz/resource"
	"phos.cc/yoo/internal/yoo/store"
)

type Biz interface {
	Resources() resource.ResourceBiz
	Menus() menu.MenuBiz
	CategoryTags() category_tag.CategoryTagBiz
}

type biz struct {
	ds store.IStore
}

var _ Biz = (*biz)(nil)

// NewBiz returns a new biz.
func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

func (b *biz) Resources() resource.ResourceBiz {
	return resource.New(b.ds)
}

func (b *biz) Menus() menu.MenuBiz {
	return menu.New(b.ds)
}

func (b *biz) CategoryTags() category_tag.CategoryTagBiz {
	return category_tag.New(b.ds)
}
