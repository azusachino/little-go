package basic

import (
	"fmt"
	"reflect"
	"runtime"
)

// 可以返回多个值
func Div(a, b int) (int, int) {
	return a / b, a % b
}

func Apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %s with args (%d, %d) \n", opName, a, b)
	return op(a, b)
}

func Sum2(nums ...int) int {
	s := 0
	for i := range nums {
		s += i
	}
	return s
}
