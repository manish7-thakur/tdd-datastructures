package avl_tree

import (
	"testing"
)

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
