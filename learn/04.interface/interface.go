package main

import "fmt"

type Duck interface {
	Traverse()
}

type YellowDuck struct {
	Food string
}

// 实现接口的方法即可
func (d YellowDuck) Traverse() {
	fmt.Println(d.Food)
}
