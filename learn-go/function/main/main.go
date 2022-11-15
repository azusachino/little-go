package main

import (
	"fmt"
	. "github.com/azusachino/little-go/learn-go/function"
)

func main() {
	a := Adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d= %d \n", i, a(i))
	}
	b := Fib()
	b()
}
