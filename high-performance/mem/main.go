package main

import (
	"fmt"
	"unsafe"
)

// 内存对齐

// Args 16
type Args struct {
	num1 int
	num2 int
}

// Flag 8
type Flag struct {
	num1 int16
	num2 int32
}

// A 16
type A struct {
	a int
	b struct{}
}

// B 8
type B struct {
	b struct{}
	a int
}

type demo1 struct {
	a int8
	b int16
	c int32
}

type demo2 struct {
	a int8
	c int32
	b int16
}

func main() {
	fmt.Println(unsafe.Sizeof(Args{}))
	fmt.Println(unsafe.Sizeof(Flag{}))
	fmt.Println(unsafe.Sizeof(A{}))
	fmt.Println(unsafe.Sizeof(B{}))
	fmt.Println(unsafe.Sizeof(demo1{})) // 8
	fmt.Println(unsafe.Sizeof(demo2{})) // 12
}
