package resource

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"phos.cc/yoo/internal/pkg/core"
	"phos.cc/yoo/internal/pkg/errno"
	veldt "phos.cc/yoo/internal/pkg/validator"
	v1 "phos.cc/yoo/pkg/api/yoo/v1"
)

func (ctrl *ResourceController) List(c *gin.Context) {
	var r v1.ListResourceRequest

	if err := c.ShouldBindQuery(&r); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(veldt.Translate(errs)), nil)
		} else {
			core.WriteResponse(c, errno.ErrBind, nil)
		}
		return
	}

	list, total, err := ctrl.b.Resources().List(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, gin.H{"data": gin.H{
		"content": list,
		"total":   total,
	},
		"code": 0})
}

func (ctrl *ResourceController) Get(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter, nil)
		return
	}

	resp, err := ctrl.b.Resources().Get(c, int32(id))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, gin.H{
		"data": resp,
		"code": 0,
	})
}

func (ctrl *ResourceController) Used(c *gin.Context) {
	var r v1.ListUsedResourceRequest

	var tags []string

	tagQuery := c.Query("tags")
	if tagQuery != "" {
		tags = strings.Split(tagQuery, ",")
	}

	r.Tags = tags

	if err := c.ShouldBindQuery(&r); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(veldt.Translate(errs)), nil)
		} else {
			core.WriteResponse(c, errno.ErrBind, nil)
		}
		return
	}

	list, letters, err := ctrl.b.Resources().Used(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, gin.H{
		"data": gin.H{
			"resources": list,
			"letters":   letters,
		},
		"code": 0})
}
