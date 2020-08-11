package main

import (
	"fmt"
	"log"
	gee_cache "my_project/cache/day_four_consistent_hash"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	gee_cache.NewGroup("scores", 2<<10, gee_cache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := gee_cache.NewHttpPool(addr)
	log.Println("gee cache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
