package avl_tree

import (
	"testing"
)

func TestInsertRoot(t *testing.T) {
	var root *Node
	item := 5
	root = insert(root, item)
	actual := root.value
	if actual != item {
		t.Errorf("Expected %d but found %d", item, actual)
	}
}

func TestInsertRootLeft(t *testing.T) {
	var root *Node
	item := 3
	root = insert(root, 5)
	root = insert(root, item)
	actual := root.left.value
	if actual != item {
		t.Errorf("Expected %d but found %d", item, actual)
	}
}

func TestInsertRootRight(t *testing.T) {
	var root *Node
	item := 6
	root = insert(root, 5)
	root = insert(root, item)
	actual := root.right.value
	if actual != item {
		t.Errorf("Expected %d but found %d", item, actual)
	}
}

func TestInsertRootRightLeft(t *testing.T) {
	var root *Node
	item := 6
	root = insert(root, 5)
	root = insert(root, 7)
	root = insert(root, item)
	actual := root.right.left.value
	if actual != item {
		t.Errorf("Expected %d but found %d", item, actual)
	}
}

func TestInsertRootLeftRight(t *testing.T) {
	var root *Node
	item := 6
	root = insert(root, 7)
	root = insert(root, 4)
	root = insert(root, item)
	actual := root.left.right.value
	if actual != item {
		t.Errorf("Expected %d but found %d", item, actual)
	}
}

func TestInsertMulti(t *testing.T) {
	var root *Node
	items := []int{3, 48, 28, 73, 62, 35}
	root = insertM(root, &items)
	actual := root.right.right.value
	if actual != items[3] {
		t.Errorf("Expected %d but found %d", items[3], actual)
	}
}

func TestPreorderTraversalEmpty(t *testing.T) {
	var root *Node
	var res = make([]int, 0, 0)
	preorder(root, &res)
	if len(res) != 0 {
		t.Errorf("Expected empty but found %v", res)
	}
}

func TestPreorderTraversalMulti(t *testing.T) {
	var root *Node
	var items = []int{2, 82, 492, 39, 32}
	root = insertM(root, &items)
	var res = make([]int, 0, 5)
	preorder(root, &res)
	var actual [5]int
	copy(actual[:], res)
	expected := [5]int{2, 82, 39, 32, 492}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}
