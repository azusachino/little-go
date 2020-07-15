package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// factory func
func createTreeNode(value int) *Node {
	// return memory address of the local variable
	return &Node{Value: value}
}

func (node Node) print() {
	fmt.Println(node.Value)
}

func (node *Node) setValue(n int) {
	node.Value = n
}

// iterate the tree
func (node *Node) traverse() {
	if node == nil {
		return
	}
	node.Left.traverse()
	node.print()
	node.Right.traverse()

}

func n() {
	var root Node
	fmt.Println(root)
	root = Node{Value: 3}
	root.Left = &Node{5, nil, nil}
	root.Right = &Node{}
	root.Right.Left = new(Node)
	fmt.Println(root)
	root.print()

	root.Right.Left.setValue(23)
	nodes := []Node{
		{Value: 2},
		{},
		{6, nil, nil},
	}
	fmt.Println(nodes)
}
