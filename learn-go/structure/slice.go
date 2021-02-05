package structure

import "fmt"

func S1() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	s := arr[2:4] // this is a slice
	print(s)
}

func UpdateSlice(s []int) {
	s[0] = 100
}

func S2() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1) // [2,3,4,5]
	fmt.Println(s2) // [5,6]
}

func S3() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1)                                                           // [2,3,4,5]
	fmt.Println(s2)                                                           // [5,6]
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1)) //s1 = [2 3 4 5], len(s1) = 4, cap(s1) = 6
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n", s2, len(s2), cap(s2)) // s2 = [5 6], len(s2) = 2, cap(s2) = 3
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	// s4 & s5 no longer belong to arr (view)
	fmt.Println(s3, s4, s5)    // [5 6 10] [5 6 10 11] [5 6 10 11 12]
	fmt.Println("arr = ", arr) // arr =  [0 1 2 3 4 5 6 10]
}

func S4() {
	var s []int
	for i := 0; i < 100; i++ {
		PrintSlice(s)
		s = append(s, i)
	}
	fmt.Println(s)

	s1 := make([]int, 16)
	PrintSlice(s1)
	copy(s1, s)

}

func PrintSlice(s []int) {
	fmt.Printf("len = %d, cap = %d \n", len(s), cap(s))
}
