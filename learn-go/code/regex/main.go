package main

import (
	"fmt"
	"regexp"
)

const text = "my email is azusachino@yahoo.com"

func main() {
	regex(text)
}

func regex(str string) {
	// re, err := regexp.Compile("^[a-z]+@(.*?)$")
	// if err != nil {
	//		panic(err)
	//	}
	//return re.FindString(str)
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)`)
	Print(re.FindAllStringSubmatch(str, -1)) // -1 means return all matches
}

func Print(obj interface{}) {
	fmt.Println(obj)
}
