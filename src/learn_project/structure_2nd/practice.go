package structure_2nd

import (
	"fmt"
)

func LengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if last, ok := lastOccurred[ch]; ok && last >= start {
			start = last + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func Acsii(s string) {
	// chinese contains three bytes
	for _, b := range []byte(s) {
		fmt.Printf("%X", b)
	}
	fmt.Println(s)

	// utf-8
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c)", i, ch)
	}
}

func Str(s string) {
	// Fields, Split, Join
	// Contains, Index
	// ToLower
	// Trim, TrimRight, TrimLeft
}
