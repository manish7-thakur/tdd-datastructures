package bst_tree

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

func postorder(node *Node, res *[]int) []int {
	if node == nil {
		return *res
	} else {
		postorder(node.left, res)
		postorder(node.right, res)
		*res = append(*res, node.value)
	}
	return *res
}

func deleteNode(node *Node, i int, parent *Node) *Node {
	if node == nil {
		return node
	} else if i < node.value {
		deleteNode(node.left, i, node)
	} else if i > node.value {
		deleteNode(node.right, i, node)
	} else {
		if parent.value == -2635 {
			node = nil
		} else if node.left == nil && node.right == nil {
			if node == parent.left {
				parent.left = nil
			} else {
				parent.right = nil
			}
		} else if node.left == nil {
			if node == parent.left {
				parent.left = node.right
			} else {
				parent.right = node.right
			}
		} else if node.right == nil {
			if node == parent.left {
				parent.left = node.left
			} else {
				parent.right = node.left
			}
		} else {
			succ := node.right
			succParent := node
			for ; succ.left != nil; {
				succParent = succ
				succ = succ.left
			}
			node.value = succ.value
			if succParent != node {
				succParent.left = succ.right
			} else {
				node.right = succ.right //could be succParent.right = succ.right
			}
		}
	}
	return node
}
