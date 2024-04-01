package main

import (
	"fmt"
	d "github.com/azusachino/golong/learn-go/interface"
)

func main() {
	duck := d.YellowDuck{Food: "battery"}
	fmt.Println(duck)
}
