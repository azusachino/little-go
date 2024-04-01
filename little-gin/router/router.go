package router

import (
	"github.com/azusachino/golong/little-gin/router/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/upload/images", http.Dir("upload/"))

	r.POST("/auth", api.GetAuth)
	return r
}
