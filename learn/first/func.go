package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	//fmt.Println(apply(div, 5, 6))
	fmt.Println(apply(func(a, b int) int {
		return a + b
	}, 2, 3))
	q, r := div(1, 2)
	fmt.Println(q, r)
	// 下划线可以占位
	c, _ := div(2, 3)
	fmt.Println(c)
}

// 可以返回多个值
func div(a, b int) (int, int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling func %s with args (%d, %d) \n", opName, a, b)
	return op(a, b)
}

func sum2(nums ...int) int {
	s := 0
	for i := range nums {
		s += i
	}
	return s
}
