package main

import (
	"fmt"
	. "github.com/azusachino/golong/learn-go/structure"
	"unicode/utf8"
)

func main() {
	A1()
	M1()
	fmt.Println(LengthOfNonRepeatingSubStr("abcasbcs"))
	fmt.Println(LengthOfNonRepeatingSubStr("你好呀"))
	fmt.Println("Rune count:", utf8.RuneCountInString("s"))
	bytes := []byte("water")
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	var arr = [5]int{1, 2, 3, 4, 5}
	s := arr[:]
	UpdateSlice(s)
	fmt.Println(s)   // [100,2,3,4,5]
	fmt.Println(arr) // [100,2,3,4,5]
	S3()
	S4()
}
