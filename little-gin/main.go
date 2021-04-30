package main

import (
	"fmt"
	"github.com/little-go/little-gin/pkg/setting"
)

func init() {
	setting.Setup()
}

func main() {
	fmt.Println("done")
}
