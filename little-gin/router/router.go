package router

import (
	"github.com/gin-gonic/gin"
	"github.com/little-go/little-gin/router/api"
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
