package main

import "fmt"

func main() {
	p1()
	a, b := 3, 4
	fmt.Println(swap(a, b)) // 3 4
	swap2(&a, &b)
	fmt.Println(a, b) // 4 3
}

func p1() {
	var a = 2
	var pa = &a
	*pa = 3
	fmt.Println(a) // 3
}

func swap(a, b int) (int, int) {
	b, a = a, b
	return b, a
}

func swap2(a, b *int) {
	*b, *a = *a, *b
}
