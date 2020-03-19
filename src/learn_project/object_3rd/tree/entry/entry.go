package main

import (
	"fmt"
	"learn_project/object_3rd/tree"
)

func main() {
	node := tree.Node{
		Value: 0,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(node)
}
