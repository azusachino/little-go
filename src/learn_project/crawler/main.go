package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	get()
}

func get() {
	res, err := http.Get("http://www.zhenai.com/zhenhun")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("err occurred")
		return
	}
	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	printCityList(all)
	//fmt.Printf("%s \n", all)
}

func printCityList(contents []byte) {
	re, err := regexp.Compile(`<a target="_blank" href="(http://www.zhenai.com/zhenhun/[a-zA-z]+)"[^>]*>([^<]+)</a>`)
	if err != nil {
		panic(err)
	}
	res := re.FindAll(contents, -1)

	for _, m := range res {
		fmt.Printf("%s \n", m)
	}
}
