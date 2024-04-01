package main

import (
	"fmt"
	"github.com/azusachino/golong/learn-project/web/day_one_http_base/base3/gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})

	engine.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	engine.POST("/post", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, request.Method)
	})

	_ = engine.Run(":9999")
}
