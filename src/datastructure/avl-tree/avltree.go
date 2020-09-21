package avl_tree

type AvlNode struct {
	value  int
	height int
	right  *AvlNode
	left   *AvlNode
}

type AVlTree struct {
	root *AvlNode
}

func (t *AVlTree) insert(value int) {
	t.root = insert(t.root, value)
}

func (t *AVlTree) delete(i int) {
	t.root = deleteItem(t.root, i)
}

func deleteItem(node *AvlNode, i int) *AvlNode {
	if node == nil {
		return nil
	} else if node.value > i {
		node.left = deleteItem(node.left, i)
	} else if node.value < i {
		node.right = deleteItem(node.right, i)
	} else {
		if node.left == nil {
			node = node.right
		} else if node.right == nil {
			node = node.left
		} else {
			succ := findSucc(node.right)
			node.value = succ.value
			node.right = deleteItem(node.right, succ.value)
		}
		if node == nil {
			return node
		}
	}
	node.height = Max(height(node.left), height(node.right)) + 1
	bf := getBalanceFactor(node)
	return rotate(node, bf)
}

func insert(node *AvlNode, value int) *AvlNode {
	if node == nil {
		node = &AvlNode{value, 0, nil, nil}
	} else if value > node.value {
		node.right = insert(node.right, value)
	} else {
		node.left = insert(node.left, value)
	}
	node.height = Max(height(node.left), height(node.right)) + 1
	bf := getBalanceFactor(node)
	return rotate(node, bf)
}

func rotate(node *AvlNode, bf int) *AvlNode {
	if bf < -1 && getBalanceFactor(node.right) <= 0 {
		return rotateLeft(node)
	} else if bf > 1 && getBalanceFactor(node.left) >= 0 { // can check value < node.left.value for insertion
		return rotateRight(node)
	} else if bf > 1 && getBalanceFactor(node.left) < 0 {
		node.left = rotateLeft(node.left)
		return rotateRight(node)
	} else if bf < -1 && getBalanceFactor(node.right) > 0 {
		node.right = rotateRight(node.right)
		return rotateLeft(node)
	} else {
		return node
	}
}

func findSucc(start *AvlNode) *AvlNode {
	for ; start.left != nil; start = start.left {

	}
	return start
}

func rotateLeft(pivot *AvlNode) *AvlNode {
	newTop := pivot.right
	pivot.right = newTop.left
	newTop.left = pivot

	pivot.height = Max(height(pivot.left), height(pivot.right)) + 1
	newTop.height = Max(height(newTop.left), height(newTop.right)) + 1
	return newTop
}

func rotateRight(pivot *AvlNode) *AvlNode {
	newTop := pivot.left
	pivot.left = newTop.right
	newTop.right = pivot

	pivot.height = Max(height(pivot.left), height(pivot.right)) + 1
	newTop.height = Max(height(newTop.left), height(newTop.right)) + 1

	return newTop
}

func getBalanceFactor(node *AvlNode) int {
	return height(node.left) - height(node.right)
}

func height(node *AvlNode) int {
	if node != nil {
		return node.height
	}
	return -1
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
