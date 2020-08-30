package avl_tree

import "testing"

type AVlTree struct {
	root *AvlNode
}

func (t *AVlTree) insert(value int) {
	t.root = insert(t.root, value)
}

func insert(root *AvlNode, value int) *AvlNode {
	if root == nil {
		root = &AvlNode{value, 0, nil}
	} else if value > root.value {
		root.right = insert(root.right, value)
		root.height = root.right.height + 1
	}
	return root
}

type AvlNode struct {
	value  int
	height int
	right  *AvlNode
}

func TestInsertSingleNode(t *testing.T) {
	tree := AVlTree{}
	tree.insert(5)
	root := tree.root
	if root.height != 0 {
		t.Errorf("Expected balance factor to be %d but found %d", 0, root.height)
	}
}

func TestInsertTwoNodes(t *testing.T) {
	tree := AVlTree{}
	tree.insert(5)
	tree.insert(7)
	root := tree.root
	if root.height != 1 || root.right.height != 0 {
		t.Errorf("Expected balance factor to be %d but found %d", -1, root.height)
	}
}
