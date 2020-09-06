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

func TestDeleteExistentItemWithOneLeftChild(t *testing.T) {
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

func TestDeleteExistentItemWithOneRightChild(t *testing.T) {
	tree := AVlTree{}
	tree.insert(7)
	tree.insert(6)
	tree.insert(5)
	tree.insert(8)
	tree.delete(7)
	root := tree.root
	if root.height != 1 || root.right.value != 8 {
		t.Errorf("Expected height to be %d but found %d", 1, root.height)
	}
}

func TestDeleteExistentItemWithTwoChildren(t *testing.T) {
	tree := AVlTree{}
	tree.insert(7)
	tree.insert(8)
	tree.insert(5)
	tree.insert(6)
	tree.insert(4)
	tree.delete(5)
	root := tree.root
	if root.height != 2 || root.right.value != 8 {
		t.Errorf("Expected height to be %d but found %d", 2, root.height)
	}
}

func TestDeleteExistentItemWithTwoChildrenHavingChildren(t *testing.T) {
	tree := AVlTree{}
	tree.insert(7)
	tree.insert(8)
	tree.insert(4)
	tree.insert(6)
	tree.insert(3)
	tree.insert(2)
	tree.insert(10)
	tree.delete(7)
	root := tree.root
	if root.height != 2 || root.right.value != 8 {
		t.Errorf("Expected height to be %d but found %d", 2, root.height)
	}
}

func TestDeleteExistentItemWithLeftRotation(t *testing.T) {
	tree := AVlTree{}
	tree.insert(4)
	tree.insert(5)
	tree.insert(3)
	tree.insert(6)
	tree.delete(3)
	root := tree.root
	if root.height != 1 || root.right.value != 6 {
		t.Errorf("Expected height to be %d but found %d", 1, root.height)
	}
}

func TestDeleteExistentItemWithLeftRotationBFRightSubtreeZero(t *testing.T) {
	tree := AVlTree{}
	tree.insert(4)
	tree.insert(6)
	tree.insert(3)
	tree.insert(5)
	tree.insert(7)
	tree.delete(3)
	root := tree.root
	if root.height != 2 || root.right.value != 7 {
		t.Errorf("Expected height to be %d but found %d", 1, root.height)
	}
}

func TestDeleteExistentItemWithRightRotation(t *testing.T) {
	tree := AVlTree{}
	tree.insert(4)
	tree.insert(5)
	tree.insert(3)
	tree.insert(2)
	tree.delete(5)
	root := tree.root
	if root.height != 1 || root.right.value != 4 {
		t.Errorf("Expected height to be %d but found %d", 1, root.height)
	}
}

func TestDeleteExistentItemWithLeftRightRotation(t *testing.T) {
	tree := AVlTree{}
	tree.insert(4)
	tree.insert(5)
	tree.insert(2)
	tree.insert(3)
	tree.delete(5)
	root := tree.root
	if root.height != 1 || root.right.value != 4 {
		t.Errorf("Expected height to be %d but found %d", 1, root.height)
	}
}

func TestDeleteExistentItemWithRightLeftRotation(t *testing.T) {
	tree := AVlTree{}
	tree.insert(3)
	tree.insert(5)
	tree.insert(2)
	tree.insert(4)
	tree.delete(2)
	root := tree.root
	if root.height != 1 || root.right.value != 5 {
		t.Errorf("Expected height to be %d but found %d", 1, root.height)
	}
}
