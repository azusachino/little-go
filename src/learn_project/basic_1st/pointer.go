package basic_1st

import "fmt"

func P1() {
	var a = 2
	var pa = &a
	*pa = 3
	fmt.Println(a) // 3
}

func Swap(a, b int) (int, int) {
	b, a = a, b
	return b, a
}

func Swap2(a, b *int) {
	*b, *a = *a, *b
}
