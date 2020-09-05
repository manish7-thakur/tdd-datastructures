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

func TestInsertThreeNodesWithRightRotation(t *testing.T) {
	tree := AVlTree{}
	tree.insert(7)
	tree.insert(6)
	tree.insert(5)
	root := tree.root
	if root.height != 1 || root.left.height != 0 {
		t.Errorf("Expected height to be %d but found %d", 1, root.height)
	}
}

func TestInsertThreeNodesWithLeftRotation(t *testing.T) {
	tree := AVlTree{}
	tree.insert(5)
	tree.insert(6)
	tree.insert(7)
	root := tree.root
	if root.height != 1 || root.left.height != 0 {
		t.Errorf("Expected height to be %d but found %d", 1, root.height)
	}
}

func TestInsertThreeNodesWithLeftRightRotation(t *testing.T) {
	tree := AVlTree{}
	tree.insert(7)
	tree.insert(5)
	tree.insert(6)
	root := tree.root
	if root.height != 1 || root.left.height != 0 {
		t.Errorf("Expected height to be %d but found %d", 1, root.height)
	}
}

func TestInsertThreeNodesWithRightLeftRotation(t *testing.T) {
	tree := AVlTree{}
	tree.insert(5)
	tree.insert(7)
	tree.insert(6)
	root := tree.root
	if root.height != 1 || root.left.height != 0 {
		t.Errorf("Expected height to be %d but found %d", 1, root.height)
	}
}

func TestDeleteNonExistentItem(t *testing.T) {
	tree := AVlTree{}
	tree.insert(7)
	tree.delete(5)
	root := tree.root
	if root.height != 0 {
		t.Errorf("Expected height to be %d but found %d", 0, root.height)
	}
}

func TestDeleteExistentItemNoChildren(t *testing.T) {
	tree := AVlTree{}
	tree.insert(7)
	tree.insert(5)
	tree.delete(5)
	root := tree.root
	if root.height != 0 {
		t.Errorf("Expected height to be %d but found %d", 0, root.height)
	}
}

func TestDeleteExistentItemOneChild(t *testing.T) {
	tree := AVlTree{}
	tree.insert(7)
	tree.insert(6)
	tree.insert(5)
	tree.insert(3)
	tree.delete(5)
	root := tree.root
	if root.height != 1 || root.left.value != 3 {
		t.Errorf("Expected height to be %d but found %d", 0, root.height)
	}
}
