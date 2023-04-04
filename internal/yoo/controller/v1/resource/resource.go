package resource

import (
	"phos.cc/yoo/internal/yoo/biz"
	"phos.cc/yoo/internal/yoo/store"
)

type ResourceController struct {
	b biz.Biz
}

func New(ds store.IStore) *ResourceController {
	return &ResourceController{b: biz.NewBiz(ds)}
}
