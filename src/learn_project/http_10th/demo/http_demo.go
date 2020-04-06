package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	h2()
}

func h1() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	s, e := httputil.DumpResponse(res, true)
	if e != nil {
		panic(e)
	}
	fmt.Printf("%s \n", s)
}

func h2() {
	req, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.162 Safari/537.36 Edg/80.0.361.109")

	// build a client
	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			return nil
		},
	}

	// basic
	// res, err := http.DefaultClient.Do(req)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	s, e := httputil.DumpResponse(res, true)
	if e != nil {
		panic(e)
	}
	fmt.Printf("%s \n", s)
}
