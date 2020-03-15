package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcasbcs"))
	fmt.Println(lengthOfNonRepeatingSubStr("你好呀"))
	fmt.Println("Rune count:", utf8.RuneCountInString("s"))
	bytes := []byte("water")
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
}

func lengthOfNonRepeatingSubStr(s string) int {
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

func acsii(s string) {
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

func str(s string) {
	// Fields, Split, Join
	// Contains, Index
	// ToLower
	// Trim, TrimRight, TrimLeft
}
