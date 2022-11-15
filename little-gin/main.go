package main

import (
	"fmt"
	"github.com/azusachino/little-go/little-gin/pkg/setting"
)

func init() {
	setting.Setup()
}

func main() {
	fmt.Println("done")
}
