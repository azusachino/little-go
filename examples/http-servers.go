package examples

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	_, _ = fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			_, _ = fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func HttpServer_() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	_ = http.ListenAndServe(":8090", nil)
}
