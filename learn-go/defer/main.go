package main

import "fmt"

func main() {
	defer1()
}

// Original String: abcd
// Reversed String: dcba
func defer1() {
	name := "abcd"
	fmt.Printf("Original String: %s \n", name)
	fmt.Printf("Reversed String: ")
	for _, v := range name {
		defer fmt.Printf("%c", v)
	}
}
