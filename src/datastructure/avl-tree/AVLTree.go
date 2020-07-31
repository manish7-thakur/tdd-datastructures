package avl_tree

type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(node *Node, i int) *Node {
	if node == nil {
		node = &Node{value: i}
	} else if i < node.value {
		node.left = insert(node.left, i)
	} else {
		node.right = insert(node.right, i)
	}
	return node
}
