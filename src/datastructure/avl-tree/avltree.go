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
	t.root = deleteItem(t.root, i, t.root)
}

func deleteItem(node *AvlNode, i int, parent *AvlNode) *AvlNode {
	if node == nil {
		return parent
	} else if node.value > i {
		parent = deleteItem(node.left, i, node)
	} else if node.value < i {
		parent = deleteItem(node.right, i, node)
	} else {
		if node.left == nil && node.right == nil {
			if parent.left == node {
				parent.left = nil
			} else {
				parent.right = nil
			}
		} else if node.left == nil {
			if parent.left == node {
				parent.left = node.right
			} else {
				parent.right = node.right
			}
		} else if node.right == nil {
			if parent.left == node {
				parent.left = node.left
			} else {
				parent.right = node.left
			}
		} else {
			succPar := node.right
			succ := succPar
			for ; succ.left != nil; succ = succ.left {
				succPar = succ
			}
			node.value = succ.value
			succPar.left = nil
			succPar.height = Max(height(succPar.left), height(succPar.right)) + 1
			if succPar == succ {
				node.right = succPar.right
			}
			node.height = Max(height(node.left), height(node.right)) + 1
		}
		parent.height = Max(height(parent.left), height(parent.right)) + 1
		bf := getBalanceFactor(parent)
		if bf < -1 && getBalanceFactor(parent.right) <= 0 {
			parent = rotateLeft(parent)
		} else if bf > 1 && getBalanceFactor(parent.left) >= 0 {
			parent = rotateRight(parent)
		} else if bf > 1 && getBalanceFactor(parent.left) < 0 { // ==0 is included above
			parent.left = rotateLeft(parent.left)
			parent = rotateRight(parent)
		} else if bf < -1 && getBalanceFactor(parent.right) > 0 {
			parent.right = rotateRight(parent.right)
			parent = rotateLeft(parent)
		}
	}
	return parent
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
