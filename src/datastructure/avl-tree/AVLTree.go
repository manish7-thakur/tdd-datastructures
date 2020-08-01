package avl_tree

type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(root *Node, i int) *Node {
	if root == nil {
		root = &Node{value: i}
	} else if i < root.value {
		root.left = insert(root.left, i)
	} else {
		root.right = insert(root.right, i)
	}
	return root
}

func insertM(root *Node, items *[]int) *Node {
	for _, v := range *items {
		root = insert(root, v)
	}
	return root
}

func preorder(node *Node, items *[]int) []int {
	if node == nil {
		return *items
	} else {
		*items = append(*items, node.value)
		preorder(node.left, items)
		preorder(node.right, items)
	}
	return *items
}

func inorder(node *Node, res *[]int) []int {
	if node == nil {
		return *res
	} else {
		inorder(node.left, res)
		*res = append(*res, node.value)
		inorder(node.right, res)
	}
	return *res
}
