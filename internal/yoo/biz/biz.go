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
	TX() TXBiz
}

type TXBiz interface {
	Biz
	Commit() error
	Rollback() error
}

type biz struct {
	ds  store.IStore
	tds store.ITXStore
}

// NewBiz returns a new biz.
func NewBiz(ds store.IStore) Biz {
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

func (b *biz) TX() TXBiz {
	return &biz{tds: b.ds.TX()}
}

func (b *biz) Commit() error {
	return b.tds.Commit()
}

func (b *biz) Rollback() error {
	return b.tds.Rollback()
}
