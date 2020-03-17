package tree

import (
	"fmt"
)

func (node *Node) Traverse() {
	node.TraverseFunc(func(node *Node) {
		node.print()
	})
	fmt.Println("")
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func main() {
	var root Node
	root.Traverse()
}
