package yoo

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "phos.cc/yoo/docs"
	"phos.cc/yoo/internal/pkg/core"
	"phos.cc/yoo/internal/pkg/errno"
	"phos.cc/yoo/internal/yoo/controller/v1/category_tag"
	"phos.cc/yoo/internal/yoo/controller/v1/files"
	"phos.cc/yoo/internal/yoo/controller/v1/menu"
	"phos.cc/yoo/internal/yoo/controller/v1/resource"
	"phos.cc/yoo/internal/yoo/store"
)

func installRouters(g *gin.Engine) error {

	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{

		// 创建 resources 路由分组
		rc := resource.New(store.S)
		resourcev1 := v1.Group("/resources")
		{
			resourcev1.PATCH("/:id", rc.Update)
			resourcev1.GET("/:id", rc.Get)
			resourcev1.GET("", rc.List)
			resourcev1.GET("/used", rc.Used)
		}

		mc := menu.New(store.S)
		menuv1 := v1.Group("/menus")
		{
			menuv1.POST("", mc.Create)
			menuv1.PATCH("/:id", mc.Update)
			menuv1.PATCH("/updates", mc.Updates)
			menuv1.GET("/tree", mc.Tree)
			menuv1.GET("/:id", mc.Get)
			menuv1.DELETE("/:id", mc.Delete)
		}

		ctc := category_tag.New(store.S)
		category_tagv1 := v1.Group("/category_tag")
		{
			category_tagv1.GET("/tree", ctc.Tree)
			category_tagv1.GET("/categories", ctc.Categories)
			category_tagv1.GET("/tags", ctc.Tags)
		}

		fc := files.New()
		filev1 := v1.Group("/files")

		{
			filev1.POST("/upload", fc.Upload)
			assets := viper.GetString("assets-path")
			filev1.Static("", assets)
		}

	}

	return nil
}
