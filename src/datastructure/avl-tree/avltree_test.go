package avl_tree

import "testing"

type AVlTree struct {
	root *AvlNode
}

func (t *AVlTree) insert(value int) {
	if t.root == nil {
		t.root = &AvlNode{value, 0}
	}
}

type AvlNode struct {
	value int
	bf    int
}

func TestInsertSingleNode(t *testing.T) {
	tree := AVlTree{}
	tree.insert(5)
	root := tree.root
	if root.bf != 0 {
		t.Errorf("Expected balance factor to be %d but found %d", 0, root.bf)
	}
}
