package main

import "fmt"

var mk = "wa"

var (
	l1 = ""
	l2 = "aw"
)

func main() {
	fmt.Printf(mk + l1 + l2)
	variable()
}

func namePrefix(s string) {
	if len(s) > 10 {
		fmt.Printf("your name length is over 10 \r\n")
	} else {
		fmt.Println(s)
	}
}

func variable() {
	// 新建的变量都有初始值
	var i int       // 0
	var s string    // empty
	var a, b = 1, 2 // 省略类型，编译器自动推断
	var e, f, g = "1", 2, true
	k, l := "kk", "ll" // 和var定义变量一样（语法糖）只能在函数内部使用
	fmt.Println(s, i)
	fmt.Println(a + b)
	fmt.Printf(e, f, g)
	fmt.Printf(k + l)

}
