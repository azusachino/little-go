package structure_2nd

import "fmt"

func A1() ([5]int, [3]int, [5]int, [4][5]int) {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 3, 4, 5, 6}

	var grid [4][5]int
	return arr1, arr2, arr3, grid
}

func A2() {
	var arr [10]int

	// index & value
	for i, v := range arr {
		fmt.Println(i)
		fmt.Println(v)
	}
}
