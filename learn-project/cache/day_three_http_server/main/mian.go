package main

import (
	"fmt"
	geeCache "github.com/little-go/learn-project/cache/day_three_http_server"
	"github.com/little-go/learn-project/cache/day_three_http_server/lru"
	"log"
	"net/http"
)

type String string

func (d String) Len() int {
	return len(d)
}

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	lruCache := lru.New(10000, nil)
	s := new(String)
	lruCache.Add("az", s)
}

func serve() {
	geeCache.NewGroup("scores", 2<<10, geeCache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := geeCache.NewHttpPool(addr)
	log.Println("gee cache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))

}
