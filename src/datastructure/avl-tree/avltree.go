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
	if bf > 1 && value < node.left.value {
		node = rotateRight(node)
	} else if bf > 1 && value > node.left.value {
		node.left = rotateLeft(node.left)
		node = rotateRight(node)
	} else if bf < -1 && value > node.right.value {
		node = rotateLeft(node)
	} else if bf < -1 && value < node.right.value {
		node.right = rotateRight(node.right)
		node = rotateLeft(node)
	}
	return node
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
