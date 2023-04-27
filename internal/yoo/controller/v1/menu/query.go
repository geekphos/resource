package menu

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"phos.cc/yoo/internal/pkg/core"
	"phos.cc/yoo/internal/pkg/errno"
	veldt "phos.cc/yoo/internal/pkg/validator"
	v1 "phos.cc/yoo/pkg/api/yoo/v1"
)

func (ctrl *MenuController) Get(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter, nil)
		return
	}

	resp, err := ctrl.b.Menus().Get(c, int32(id))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}

func (ctrl *MenuController) Tree(c *gin.Context) {
	var r v1.ListMenuRequest

	if err := c.ShouldBindQuery(&r); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(veldt.Translate(errs)), nil)
		} else {
			core.WriteResponse(c, errno.ErrBind, nil)
		}
		return
	}

	resp, err := ctrl.b.Menus().Tree(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, gin.H{
		"data": resp,
		"code": 0,
	})
}

func (ctrl *MenuController) GetMenuPath(c *gin.Context) {
	name := c.Param("name")

	resp, err := ctrl.b.Menus().GetMenuPath(c, name)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}

func (ctrl *MenuController) GetLeaveMenus(c *gin.Context) {
	var r v1.GetLeaveMenuRequest

	r.Categories = c.QueryArray("categories")

	if err := c.ShouldBindQuery(&r); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(veldt.Translate(errs)), nil)
		} else {
			core.WriteResponse(c, errno.ErrBind, nil)
		}
		return
	}

	resp, letters, err := ctrl.b.Menus().GetLeaveMenus(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, gin.H{
		"data": gin.H{
			"menus":   resp,
			"letters": letters,
		},
		"code": 0,
	})
}
