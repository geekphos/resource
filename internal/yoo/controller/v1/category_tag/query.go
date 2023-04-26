package category_tag

import (
	"github.com/gin-gonic/gin"

	"phos.cc/yoo/internal/pkg/core"
)

func (ctrl *CategoryTagController) Tree(c *gin.Context) {
	resp, err := ctrl.b.CategoryTags().Tree(c)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, gin.H{
		"data": resp,
		"code": 0,
	})
}

func (ctrl *CategoryTagController) Categories(c *gin.Context) {
	resp, err := ctrl.b.CategoryTags().Categories(c)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, gin.H{
		"data": resp,
		"code": 0,
	})
}

func (ctrl *CategoryTagController) Tags(c *gin.Context) {
	resp, err := ctrl.b.CategoryTags().Tags(c)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, gin.H{
		"data": resp,
		"code": 0,
	})
}
