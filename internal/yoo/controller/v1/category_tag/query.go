package category_tag

import (
	"github.com/gin-gonic/gin"

	"phos.cc/yoo/internal/pkg/core"
)

func (ctrl *CategoryTagController) All(c *gin.Context) {
	resp, err := ctrl.b.CategoryTags().All(c)
	if err != nil {
		core.WriteResponse(c, err, nil)
	}
	core.WriteResponse(c, nil, resp)
}
