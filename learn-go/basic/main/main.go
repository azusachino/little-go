package main

import (
	"fmt"
	. "github.com/azusachino/little-go/learn-go/basic"
)

func main() {
	fmt.Println(Eval(1, 2, "-"))
	Aa()
	fmt.Println(Convert2Bin(5))    // 101
	fmt.Println(Convert2Bin(13))   // 1101
	fmt.Println(Convert2Bin(2315)) // 100100001011
	fmt.Println(Convert2Bin(0))    // 0
	PrintFile("str.go")
	//fmt.Println(apply(div, 5, 6))
	fmt.Println(Apply(func(a, b int) int {
		return a + b
	}, 2, 3))
	q, r := Div(1, 2)
	fmt.Println(q, r)
	// 下划线可以占位
	c, _ := Div(2, 3)
	fmt.Println(c)
}
