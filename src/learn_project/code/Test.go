package main

import "fmt"

type Call interface {
	call(int)
}

type MyFunc func(int)

func (f MyFunc) call(i int) {
	f(i)
}

func one(i int, f func(int)) {
	two(i, MyFunc(f))
}

func two(i int, c Call) {
	c.call(i)
}

func (f MyFunc) PrintSth(i int) {
	f(i)
}

func cb(i int) {
	fmt.Printf("this is callback, the value is %d \n", i)
}

func main() {
	one(2, cb)
	m := MyFunc(func(i int) {
		fmt.Printf("the val you input is %d \n", i)
	})
	m.call(2)
	m.PrintSth(3213)
}
