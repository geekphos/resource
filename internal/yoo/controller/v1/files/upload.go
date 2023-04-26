package files

import (
	"github.com/spf13/viper"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"phos.cc/yoo/internal/pkg/core"
	"phos.cc/yoo/internal/pkg/errno"
)

func (ctrl *FileController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	filename := uuid.NewString() + filepath.Ext(file.Filename)

	assets := viper.GetString("assets-path")

	if err := c.SaveUploadedFile(file, assets+"/"+filename); err != nil {
		core.WriteResponse(c, errno.InternalServerError, nil)
		return
	}

	core.WriteResponse(c, nil, gin.H{
		"data": filename,
		"code": 0,
	})

}
