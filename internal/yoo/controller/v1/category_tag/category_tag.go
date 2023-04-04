package category_tag

import (
	"phos.cc/yoo/internal/yoo/biz"
	"phos.cc/yoo/internal/yoo/store"
)

type CategoryTagController struct {
	b biz.Biz
}

func New(ds store.IStore) *CategoryTagController {
	return &CategoryTagController{b: biz.NewBiz(ds)}
}
