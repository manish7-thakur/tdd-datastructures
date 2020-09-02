package avl_tree

import (
	"testing"
)

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
	if bf == 2 {
		node = rotateRight(node)
	}
	return node
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

func rotateRight(pivot *AvlNode) *AvlNode {
	newTop := pivot.left
	pivot.left = newTop.right
	newTop.right = pivot

	pivot.height = Max(height(pivot.left), height(pivot.right)) + 1
	newTop.height = Max(height(newTop.left), height(newTop.right)) + 1

	return newTop
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type AvlNode struct {
	value  int
	height int
	right  *AvlNode
	left   *AvlNode
}

func TestInsertSingleNode(t *testing.T) {
	tree := AVlTree{}
	tree.insert(5)
	root := tree.root
	if root.height != 0 {
		t.Errorf("Expected height to be %d but found %d", 0, root.height)
	}
}

func TestInsertTwoNodes(t *testing.T) {
	tree := AVlTree{}
	tree.insert(5)
	tree.insert(7)
	root := tree.root
	if root.height != 1 || root.right.height != 0 {
		t.Errorf("Expected height to be %d but found %d", -1, root.height)
	}
}

func TestInsertThreeNodes(t *testing.T) {
	tree := AVlTree{}
	tree.insert(5)
	tree.insert(7)
	tree.insert(4)
	root := tree.root
	if root.height != 1 || root.left.height != 0 {
		t.Errorf("Expected height to be %d but found %d", -1, root.height)
	}
}

func TestInsertThreeNodesWithRotation(t *testing.T) {
	tree := AVlTree{}
	tree.insert(7)
	tree.insert(5)
	tree.insert(4)
	root := tree.root
	if root.height != 1 || root.left.height != 0 {
		t.Errorf("Expected height to be %d but found %d", 1, root.height)
	}
}
