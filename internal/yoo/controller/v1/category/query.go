package category

import (
	"github.com/gin-gonic/gin"
	"phos.cc/yoo/internal/pkg/core"
)

func (ctrl *CategoryController) All(c *gin.Context) {

	categories, err := ctrl.b.Categories().All(c)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, gin.H{
		"data": categories,
		"code": 0,
	})
}
