package routers

import (
	_ "blog-server/docs"
	v1 "blog-server/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//url:=ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Upate)
		apiv1.PATCH("/articles/:id/state", article.Upate)
		apiv1.GET("/articles", article.List)
		apiv1.GET("/articles/:id", article.Get)
	}
	return r
}
