package router

import (
	"github.com/gin-gonic/gin"
	"hex-go/http/controller"
)

// App is
func App(router *gin.Engine) {
	version := "/v1"
	v1 := router.Group(version)
	{
		v1.GET("/articles", controller.ArticleIndex)            // 記事取得
		v1.POST("/articles/create", controller.ArticleCreate)   // 記事作成
		v1.PUT("/articles/update", controller.ArticleUpdate)    // 記事更新
		v1.DELETE("/articles/delete", controller.ArticleDelete) // 記事削除
	}
}
