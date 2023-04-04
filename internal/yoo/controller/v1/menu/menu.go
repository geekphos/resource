package menu

import (
	"phos.cc/yoo/internal/yoo/biz"
	"phos.cc/yoo/internal/yoo/store"
)

type MenuController struct {
	b biz.Biz
}

func New(ds store.IStore) *MenuController {
	return &MenuController{b: biz.NewBiz(ds)}
}
