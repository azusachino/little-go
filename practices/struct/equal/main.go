package main

import "fmt"

type Vertex struct {
	Name   string
	Gender string
}

type Value struct {
	Name   string
	Gender *string
}

type Array struct {
	Name  string
	Goods []string
}

func main() {
	t1()
}

func t1() {
	v1 := Vertex{
		Name:   "1",
		Gender: "2",
	}
	v2 := Vertex{
		Name:   "1",
		Gender: "2",
	}

	if v1 == v2 { // true
		fmt.Println("hallo")
		return
	}
	fmt.Println("hello")
}

func t2() {
	v1 := Value{
		Name:   "1",
		Gender: new(string),
	}
	v2 := Value{
		Name:   "1",
		Gender: new(string),
	}
	if v1 == v2 {
		fmt.Println("hallo")
		return
	}
	fmt.Println("hello")
}

func t3() {
	v1, v2 := Array{
		Name:  "1",
		Goods: []string{"1"},
	}, Array{
		Name:  "1",
		Goods: []string{"1"},
	}
	// IDE报错，没定义==方法
	//if v1 == v2 {
	//
	//}
	println(v1, v2)
}

func t4() {
	v1 := Value{
		Name:   "1",
		Gender: nil,
	}

	v2 := &Value{
		Name:   "2",
		Gender: nil,
	}

	v3 := new(Value)

	v1.Gender = new(string)
	v2.Gender = new(string)

	if v1 == *v2 {
	}

	if v2 == v3 {

	}
}
