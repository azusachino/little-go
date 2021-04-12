package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// interface 并不是一个指针类型, 必须得类型和值同时都为 nil 的情况下，interface 的 nil 判断才会为 true。

/**
  interface共有两类数据类型
  runtime.eface 结构体：表示不包含任何方法的空接口，也称为 empty interface。
  runtime.iface 结构体：表示包含方法的接口。
*/
type _type struct{}

type itab struct{}

type eface struct {
	_type *_type
	data  unsafe.Pointer
}

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

func main() {
	var data *byte
	var in interface{}

	in = data
	fmt.Println(IsNil(in))
}

func m1() {
	// false

	var v interface{}
	v = (*int)(nil)
	fmt.Println(v == nil)
}

func m2() {
	//<nil> true
	//<nil> true
	//<nil> false

	var data *byte
	var in interface{}

	fmt.Println(data, data == nil)
	fmt.Println(in, in == nil)

	in = data
	fmt.Println(in, in == nil)
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
