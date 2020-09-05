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
	deleteItem(t.root, i, nil)
}

func deleteItem(node *AvlNode, i int, parent *AvlNode) {
	if node == nil {
		return
	} else if node.value > i {
		deleteItem(node.left, i, node)
	} else if node.value < i {
		deleteItem(node.right, i, node)
	} else {
		if parent.left == node {
			parent.left = findReplacement(node)
		} else {
			parent.right = findReplacement(node)
		}
		parent.height = Max(height(parent.left), height(parent.right)) + 1
	}
}

func findReplacement(node *AvlNode) *AvlNode {
	if node.left == nil {
		return node.right
	} else if node.right == nil {
		return node.left
	}
	return nil
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
		return rotateRight(node)
	} else if bf > 1 && value > node.left.value {
		node.left = rotateLeft(node.left)
		return rotateRight(node)
	} else if bf < -1 && value > node.right.value {
		return rotateLeft(node)
	} else if bf < -1 && value < node.right.value {
		node.right = rotateRight(node.right)
		return rotateLeft(node)
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
