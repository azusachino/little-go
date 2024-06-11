package examples

import (
	"fmt"
	"sort"
)

func Sort__() {
	strings := []string{"c", "b", "a"}
	sort.Strings(strings)
	println(strings)

	nums := []int{4, 5, 7, 8}
	sort.Ints(nums)
	println(nums)

	s := sort.IntsAreSorted(nums)
	fmt.Println("Sorted: ", s)
}
