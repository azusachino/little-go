package tree

func (node Node) getValue() int {
	return node.Value
}

// 别名实现
type MyTreeNode struct {
	node *Node
}

func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := MyTreeNode{myNode.node.Left}
	left.postOrder()
	right := MyTreeNode{myNode.node.Right}
	right.postOrder()
	// MyTreeNode{myNode.node.Left}.postOrder()
	// MyTreeNode{myNode.node.Right}.postOrder()
	myNode.node.print()
}
