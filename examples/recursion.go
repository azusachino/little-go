package examples

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func Recursion_() {
	fmt.Println(fact(7))
}
