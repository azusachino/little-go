package main

import "fmt"

func main() {
	order2()
}

/**
5
4
3
2
1
0
*/
func order() {
	//在 for 循环时，局部变量 i 已经传入进 defer func 中 ，属于值传递。其值在 defer 语句声明时的时候就已经确定下来了。
	//结合 defer 关键字的特性，是按先进后出的顺序来执行的。
	var whatever [6]struct{}
	for i := range whatever {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

/**
5
5
5
5
5
5
*/
func order2() {
	//在 for 循环结束后，局部变量 i 的值已经是 5 了，并且 defer的闭包是直接引用变量的 i。
	//结合defer 关键字的特性，可得知会在 main 方法主体结束后再执行。
	var whatever [6]struct{}
	for i := range whatever {
		defer func() {
			fmt.Println(i)
		}()
	}
}
