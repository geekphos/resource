package resource

import (
	"context"

	"github.com/jinzhu/copier"
	"gorm.io/datatypes"

	"phos.cc/yoo/internal/pkg/errno"
	"phos.cc/yoo/internal/pkg/model"
	"phos.cc/yoo/internal/yoo/store"
	v1 "phos.cc/yoo/pkg/api/yoo/v1"
)

type ResourceBiz interface {
	Update(ctx context.Context, r *v1.UpdateResourceRequest) error
	List(ctx context.Context, r *v1.ListResourceRequest) ([]*v1.ListResourceResponse, int64, error)
	Get(ctx context.Context, id int32) (*v1.GetResourceResponse, error)
}

type resourceBiz struct {
	ds store.IStore
}

var _ ResourceBiz = (*resourceBiz)(nil)

func New(ds store.IStore) *resourceBiz {
	return &resourceBiz{ds: ds}
}

func (b *resourceBiz) Update(ctx context.Context, r *v1.UpdateResourceRequest) error {
	resourceM, err := b.ds.Resources().Get(ctx, r.ID)
	if err != nil {
		return errno.InternalServerError
	}
	_ = copier.Copy(resourceM, r)

	if err := b.ds.Resources().Update(ctx, resourceM); err != nil {
		return errno.InternalServerError
	}
	return nil
}

func (b *resourceBiz) List(ctx context.Context, r *v1.ListResourceRequest) ([]*v1.ListResourceResponse, int64, error) {
	var resourceM = &model.ResourceM{}
	_ = copier.Copy(resourceM, r)

	if r.Tag != "" {
		resourceM.Tags = datatypes.JSON([]byte(`["` + r.Tag + `"]`))
	}

	resources, total, err := b.ds.Resources().List(ctx, r.Page, r.PageSize, resourceM)
	if err != nil {
		return nil, 0, errno.InternalServerError
	}

	var resourcesR []*v1.ListResourceResponse
	for _, resource := range resources {
		var resourceR = &v1.ListResourceResponse{}
		_ = copier.Copy(resourceR, resource)
		resourcesR = append(resourcesR, resourceR)
	}

	return resourcesR, total, nil
}

func (b *resourceBiz) Get(ctx context.Context, id int32) (*v1.GetResourceResponse, error) {
	resource, err := b.ds.Resources().Get(ctx, id)
	if err != nil {
		return nil, errno.ErrResourceNotFound
	}

	var resourceR = &v1.GetResourceResponse{}
	_ = copier.Copy(resourceR, resource)

	return resourceR, nil
}