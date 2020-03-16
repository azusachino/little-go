package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
	"time"
)

const url = "https://www.baidu.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, data map[string]string) string
}

// composite interface
type RP interface {
	Retriever
	Poster
}

func download(r Retriever, url string) string {
	return r.Get(url)
}

func Post(p Poster, url string) string {
	return p.Post(url, map[string]string{
		"name": "c",
		"age":  "23",
	})
}

func session(s RP) string {
	s.Post(url, map[string]string{
		"contents": "another fake",
	})
	return s.Get(url)
}

func main() {
	var rr Retriever
	// switch (type)
	inspect(rr)
	var r = mock.Retriever{Contents: "this is test message"}
	var r2 = real.Retriever{}
	var r3 = real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Println(download(&r, "http://bing.com"))
	fmt.Println(r2.Get("https://www.cn.bing.com"))
	fmt.Println(r3.Get(""))

	// type assertion
	realR := rr.(*real.Retriever)
	fmt.Println(realR.TimeOut)

	if fakeR, ok := rr.(mock.Retriever); ok {
		fmt.Println(fakeR)
	}
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	case mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	default:
		fmt.Println("none retriever recognised!!!")
	}
}
