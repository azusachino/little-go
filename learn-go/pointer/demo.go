package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(&a, b) // 两个都是指针地址

	fmt.Println(a, *b) // 两个都是数值

	a_()
}

func a_() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p \n", a, b, c)  // b,c都是指针
	fmt.Printf("%v %v %p \n", a, *b, c) // b,c都是指针
}
