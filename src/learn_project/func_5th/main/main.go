package main

import (
	"fmt"
	"learn_project/func_5th"
)

func main() {
	a := func_5th.Adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d= %d \n", i, a(i))
	}
	b := func_5th.Fib()
	b()
}
