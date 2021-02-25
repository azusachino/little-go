package main

import (
	"fmt"
	"log"
	"net/http"
)

// 实现Handler
type Engine struct {
}

func (engine *Engine) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		_, _ = fmt.Fprintf(rw, "Url.Path = %q \n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			_, _ = fmt.Fprintf(rw, "Header [%q] = %q \n", k, v)
		}
	default:
		_, _ = fmt.Fprintf(rw, "404 Not Found: %s \n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
