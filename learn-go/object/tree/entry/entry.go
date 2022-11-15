package main

import (
	"fmt"
	"github.com/azusachino/little-go/learn-go/object/tree"
)

func main() {
	node := tree.Node{
		Value: 0,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(node)
}
