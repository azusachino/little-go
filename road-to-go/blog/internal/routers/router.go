package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/little-go/road-to-go/blog/internal/routers/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	article := v1.NewArticle()
	tag := v1.NewTag()
	apiV1 := r.Group("/api/v1")

	{
		apiV1.POST("/tags", tag.Create)

		apiV1.POST("/articles", article.Create)
	}

	return r
}
