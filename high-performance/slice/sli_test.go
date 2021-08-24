package main

import (
	"fmt"
	"runtime"
	"testing"
)

func square(arr *[3]int) {
	for i, num := range *arr {
		(*arr)[i] = num * num
	}
}

func mp(m *map[int]interface{}) {
	if m != nil {
		for k, v := range *m {
			fmt.Println(k, v)
		}
	}
}

func printLenCap(nums []int) {
	fmt.Printf("len: %d, cap: %d %v\n", len(nums), cap(nums), nums)
}

func testLastChars(t *testing.T, f func([]int) []int) {
	t.Helper()
	ans := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := generateWithCap(128 * 1024) // 1M
		ans = append(ans, f(origin))
		runtime.GC()
	}
	printMem(t)
	_ = ans
}

func TestLastCharsBySlice(t *testing.T) { testLastChars(t, lastNumsBySlice) }
func TestLastCharsByCopy(t *testing.T)  { testLastChars(t, lastNumsByCopy) }

func printMem(t *testing.T) {
	t.Helper()
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	t.Logf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
}

func TestSliceLenAndCap(t *testing.T) {
	nums := []int{1}
	printLenCap(nums) // len: 1, cap: 1 [1]
	nums = append(nums, 2)
	printLenCap(nums) // len: 2, cap: 2 [1 2]
	nums = append(nums, 3)
	printLenCap(nums) // len: 3, cap: 4 [1 2 3]
	nums = append(nums, 3)
	printLenCap(nums) // len: 4, cap: 4 [1 2 3 3]
}

func TestMap(t *testing.T) {
	m := make(map[int]interface{})
	m[1] = 2
	mp(&m)
}

func TestArrayPointer(t *testing.T) {
	a := [...]int{1, 2, 3} //array
	b := []int{1, 2, 3}    // slice
	square(&a)
	fmt.Println(a, b) // [1 4 9]
	if a[1] != 4 && a[2] != 9 {
		t.Fatal("failed")
	}
}
