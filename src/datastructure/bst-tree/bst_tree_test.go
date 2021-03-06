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

func TestPreorderTraversalIterEmpty(t *testing.T) {
	var root *Node
	var res = make([]int, 0, 0)
	res = PreorderIter(root)
	if len(res) != 0 {
		t.Errorf("Expected empty but found %v", res)
	}
	root = insert(root, 1)
	res = PreorderIter(root)
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{1}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}

}

func TestPreorderTraversalIterTwoItems(t *testing.T) {
	var root *Node
	var res = make([]int, 0, 0)
	items := []int{3, 1}
	root = insertM(root, &items)
	res = PreorderIter(root)
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{3, 1}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestPreorderTraversalIterThreeItems(t *testing.T) {
	var root *Node
	var res = make([]int, 0, 0)
	items := []int{3, 1, 2}
	root = insertM(root, &items)
	res = PreorderIter(root)
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{3, 1, 2}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestPreorderTraversalIterFourItems(t *testing.T) {
	var root *Node
	var res = make([]int, 0, 0)
	items := []int{3, 1, 2, 4}
	root = insertM(root, &items)
	res = PreorderIter(root)
	actual := [4]int{}
	copy(actual[:], res)
	expected := [4]int{3, 1, 2, 4}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestPreorderTraversalIterMulti(t *testing.T) {
	var root *Node
	var items = []int{2, 82, 492, 39, 32}
	root = insertM(root, &items)
	var res = PreorderIter(root)
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
	deleteNode(root, 4)
	var actual []int
	postorder(root, &actual)
	if len(actual) != 0 {
		t.Errorf("Expected empty but found %v", actual)
	}
}

func TestDeleteRootNodeNoChild(t *testing.T) {
	var root *Node
	root = insert(root, 2)
	root = deleteNode(root, 2) // special value to delete node
	if root != nil {
		t.Errorf("Root was not deleted, found %v", root)
	}
}

func TestDeleteRootLeftLeafNoChild(t *testing.T) {
	var root *Node
	root = insertM(root, &[]int{3, 28, 2})
	root = deleteNode(root, 2)
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
	root = deleteNode(root, 28)
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
	root = deleteNode(root, 1)
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
	root = deleteNode(root, 29)
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
	root = deleteNode(root, 2)
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
	root = deleteNode(root, 25)
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
	root = deleteNode(root, 30)
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
	root = deleteNode(root, 30)
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
	root = deleteNode(root, 2)
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
	root = deleteNode(root, 2)
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
	root = deleteNode(root, 20)
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
	root = deleteNode(root, 2)
	var res []int
	inorder(root, &res)
	var actual [8]int
	copy(actual[:], res)
	expected := [8]int{1, 3, 5, 8, 9, 20, 21, 36}
	if actual != expected {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}
