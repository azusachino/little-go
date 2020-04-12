package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 5, 7, 1, 2, 4}
	sort.Ints(a)
	for _, v := range a {
		fmt.Println(v)
	}
}
