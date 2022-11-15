package main

import (
	"github.com/azusachino/little-go/learn-project/web/day_two_context/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.Param{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	_ = r.Run(":9999")
}
