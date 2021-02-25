package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":9999", nil))
}

// 实现HandleFunc

func indexHandler(rw http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(rw, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		_, _ = fmt.Fprintf(rw, "Header[%q] = %q\n", k, v)
	}
}
