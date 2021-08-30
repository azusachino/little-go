package main

import (
	"github.com/little-go/road-to-go/blog/internal/routers"
	"net/http"
	"time"
)

func main() {
	r := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
