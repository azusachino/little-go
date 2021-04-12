package main

import "fmt"

type User struct {
	Id     uint64
	Name   string
	Avatar string
}

func main() {
	_ = GetUserInfo()
}

func GetUserInfo() *User {
	return &User{}
}

func o() {
	// 分配到栈上
	str := new(string)
	*str = "ASDAS"
}

//当形参为 interface 类型时，在编译阶段编译器无法确定其具体的类型。因此会产生逃逸，最终分配到堆上。
func o2() {
	// 分配到堆上
	str := new(string)
	*str = "EDDYCJY"

	fmt.Println(str)
}

//我们注意到 leaking param 的表述，它说明了变量 u 是一个泄露参数。
//结合代码可得知其传给 GetUserInfo 方法后，没有做任何引用之类的涉及变量的动作，直接就把这个变量返回出去了。
//因此这个变量实际上并没有逃逸，它的作用域还在 main() 之中，所以分配在栈上。
func _GetUserInfo(u *User) *User {
	return u
}

// 分配到堆上
func __GetUserInfo(u User) *User {
	return &u
}
