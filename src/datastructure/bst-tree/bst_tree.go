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

func deleteNode(node *Node, i int) *Node {
	if node == nil {
		return node
	} else if i < node.value {
		node.left = deleteNode(node.left, i)
	} else if i > node.value {
		node.right = deleteNode(node.right, i)
	} else {
		if node.left == nil {
			node = node.right
		} else if node.right == nil {
			node = node.left
		} else {
			succ := findSucc(node.right)
			node.value = succ.value
			node.right = deleteNode(node.right, succ.value)
		}
	}
	return node
}

func findSucc(start *Node) *Node {
	for ; start.left != nil; start = start.left {
	}
	return start
}
