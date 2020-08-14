package bst_tree

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

func TestInorderTraversalMulti(t *testing.T) {
	var root *Node
	var items = []int{2, 82, 492, 39, 32}
	var res = make([]int, 0, 5)
	root = insertM(root, &items)
	inorder(root, &res)
	var actual [5]int
	copy(actual[:], res)
	expected := [5]int{2, 32, 39, 82, 492}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestPostorderTraversalMulti(t *testing.T) {
	var root *Node
	var items = []int{2, 82, 492, 39, 32}
	var res = make([]int, 0, 5)
	root = insertM(root, &items)
	postorder(root, &res)
	var actual [5]int
	copy(actual[:], res)
	expected := [5]int{32, 39, 492, 82, 2}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestDeleteEmpty(t *testing.T) {
	var root *Node
	root = deleteNode(root, 4, &Node{right: root})
	var actual []int
	postorder(root, &actual)
	if len(actual) != 0 {
		t.Errorf("Expected empty but found %v", actual)
	}
}

func TestDeleteRootNodeNoChild(t *testing.T) {
	var root *Node
	root = insert(root, 2)
	root = deleteNode(root, 2, &Node{right: root, value: -2635}) // special value to delete node
	if root != nil {
		t.Errorf("Root was not deleted, found %v", root)
	}
}

func TestDeleteRootLeftLeafNoChild(t *testing.T) {
	var root *Node
	root = insertM(root, &[]int{3, 28, 2})
	root = deleteNode(root, 2, &Node{right: root})
	var res []int
	inorder(root, &res)
	var actual [2]int
	copy(actual[:], res)
	expected := [2]int{3, 28}
	if actual != expected {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}

func TestDeleteRootRightLeafNoChild(t *testing.T) {
	var root *Node
	root = insertM(root, &[]int{3, 28, 2})
	root = deleteNode(root, 28, &Node{right: root})
	var res []int
	postorder(root, &res)
	var actual [2]int
	copy(actual[:], res)
	expected := [2]int{2, 3}
	if actual != expected {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}

func TestDeleteLeftLeafNoChild(t *testing.T) {
	var root *Node
	root = insertM(root, &[]int{3, 28, 2, 37, 26, 1})
	root = deleteNode(root, 1, &Node{right: root})
	var res []int
	inorder(root, &res)
	var actual [5]int
	copy(actual[:], res)
	expected := [5]int{2, 3, 26, 28, 37}
	if actual != expected {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}

func TestDeleteRightLeafNoChild(t *testing.T) {
	var root *Node
	root = insertM(root, &[]int{3, 1, 2, 28, 29, 27})
	root = deleteNode(root, 29, &Node{right: root})
	var res []int
	postorder(root, &res)
	var actual [5]int
	copy(actual[:], res)
	expected := [5]int{2, 1, 27, 28, 3}
	if actual != expected {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}

func TestDeleteLeftNodeWithLeftChild(t *testing.T) {
	var root *Node
	root = insertM(root, &[]int{3, 2, 1, 28, 29, 27})
	root = deleteNode(root, 2, &Node{right: root})
	var res []int
	inorder(root, &res)
	var actual [5]int
	copy(actual[:], res)
	expected := [5]int{1, 3, 27, 28, 29}
	if actual != expected {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}

func TestDeleteLeftNodeWithRightChild(t *testing.T) {
	var root *Node
	root = insertM(root, &[]int{3, 2, 1, 30, 25, 27})
	root = deleteNode(root, 25, &Node{right: root})
	var res []int
	inorder(root, &res)
	var actual [5]int
	copy(actual[:], res)
	expected := [5]int{1, 2, 3, 27, 30}
	if actual != expected {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}

func TestDeleteRightNodeWithLeftChild(t *testing.T) {
	var root *Node
	root = insertM(root, &[]int{3, 2, 1, 25, 30, 27})
	root = deleteNode(root, 30, &Node{right: root})
	var res []int
	inorder(root, &res)
	var actual [5]int
	copy(actual[:], res)
	expected := [5]int{1, 2, 3, 25, 27}
	if actual != expected {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}

func TestDeleteRightNodeWithRightChild(t *testing.T) {
	var root *Node
	root = insertM(root, &[]int{3, 2, 1, 25, 30, 31})
	root = deleteNode(root, 30, &Node{right: root})
	var res []int
	inorder(root, &res)
	var actual [5]int
	copy(actual[:], res)
	expected := [5]int{1, 2, 3, 25, 31}
	if actual != expected {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}

func TestDeleteLeftNodeWithTwoChildren(t *testing.T) {
	var root *Node
	root = insertM(root, &[]int{20, 2, 8, 3, 21, 36, 1})
	root = deleteNode(root, 2, &Node{right: root})
	var res []int
	inorder(root, &res)
	var actual [6]int
	copy(actual[:], res)
	expected := [6]int{1, 3, 8, 20, 21, 36}
	if actual != expected {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}

func TestDeleteLeftNodeWithTwoChildrenNoSucc(t *testing.T) {
	var root *Node
	root = insertM(root, &[]int{20, 2, 8, 21, 36, 1})
	root = deleteNode(root, 2, &Node{right: root})
	var res []int
	inorder(root, &res)
	var actual [5]int
	copy(actual[:], res)
	expected := [5]int{1, 8, 20, 21, 36}
	if actual != expected {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}

func TestDeleteRootNodeWithTwoChildren(t *testing.T) {
	var root *Node
	root = insertM(root, &[]int{20, 2, 8, 21, 36, 1})
	root = deleteNode(root, 20, &Node{right: root})
	var res []int
	inorder(root, &res)
	var actual [5]int
	copy(actual[:], res)
	expected := [5]int{1, 2, 8, 21, 36}
	if actual != expected {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}

func TestDeleteLeftNodeWithTwoChildrenSuccWithRightChild(t *testing.T) {
	var root *Node
	root = insertM(root, &[]int{20, 2, 8, 3, 5, 9, 21, 36, 1})
	root = deleteNode(root, 2, &Node{right: root})
	var res []int
	inorder(root, &res)
	var actual [8]int
	copy(actual[:], res)
	expected := [8]int{1, 3, 5, 8, 9, 20, 21, 36}
	if actual != expected {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}
