package category

import (
	"phos.cc/yoo/internal/yoo/biz"
	"phos.cc/yoo/internal/yoo/store"
)

type CategoryController struct {
	b biz.Biz
}

func New(ds store.IStore) *CategoryController {
	return &CategoryController{b: biz.NewBiz(ds)}
}
