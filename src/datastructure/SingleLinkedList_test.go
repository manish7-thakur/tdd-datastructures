package datastructure

import (
	"testing"
)

/*func TestInsertFirst(t *testing.T) {
	var first *Node
	item := 3
	res := insert(first, item)
	if res.data != item {
		t.Errorf("The item wan't inserted in the list")
	}
}

func TestInsertTwo(t *testing.T) {
	var first *Node
	item1 := 3
	item2 := 4
	first = insert(first, item1)
	insert(first, item2)
	actual := first.next.data
	if actual != item2 {
		t.Errorf("%d is not equal to %d", actual, item2)
	}
}

func TestInsertThree(t *testing.T) {
	var first *Node
	item1 := 3
	item2 := 4
	item3 := 6
	first = insert(first, item1)
	insert(first, item2)
	insert(first, item3)
	actual := first.next.next.data
	if actual != item3 {
		t.Errorf("%d is not equal to %d", actual, item3)
	}
}

func TestInsertMulti(t *testing.T) {
	var first *Node
	elements := []int{3, 6, 7, 2, 4, 6, 3, 56, 53}
	first = insertM(first, elements)
	var actual, expected [9]int
	copy(expected[:], elements)
	copy(actual[:], getElements(first))
	if actual != expected {
		t.Errorf("%d is not equal to %d", actual, expected)
	}
}

func TestGetElementsEmpty(t *testing.T) {
	var first *Node
	res := getElements(first)
	if len(res) != 0 {
		t.Errorf("not empty")
	}
}

func TestGetElementsSingle(t *testing.T) {
	var first *Node
	first = insert(first, 2)
	res := getElements(first)
	if len(res) != 1 && res[0] != 2 {
		t.Errorf("elements not as expected")
	}
}

func TestGetElementsMulti(t *testing.T) {
	var first *Node
	first = insert(first, 2)
	insert(first, 4)
	insert(first, 5)
	res := [3]int{}
	copy(res[:], getElements(first))
	if res != [3]int{2, 4, 5} {
		t.Errorf("elements not as expected")
	}
}

func TestDeleteItemEmpty(t *testing.T) {
	var first *Node
	deleteItem(first, 3)
	res := [1]int{}
	copy(res[:], getElements(first))
	if res != [1]int{0} {
		t.Errorf("Item not deleted")
	}
}*/

/*func TestDeleteItemFirst(t *testing.T) {
	var first *Node
	first = insert(first, 2)
	first = deleteItem(first, 2)
	res := [1]int{}
	copy(res[:], getElements(first))
	expected := [1]int{0}
	if res != expected {
		t.Errorf("Expected %v but was %v", expected, res)
	}
}*/

func TestDeleteItemMiddle(t *testing.T) {
	var first *Node
	elements := []int{3, 6, 7, 2, 4, 6, 3, 56, 53}
	first = insertM(first, elements)
	first = deleteItem(first, 4)
	var actual [8]int
	var expected = [8]int{3, 6, 7, 2, 6, 3, 56, 53}
	copy(actual[:], getElements(first))
	if actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestDeleteItemLast(t *testing.T) {
	var first *Node
	elements := []int{3, 6, 7, 2, 6, 3, 56, 53}
	first = insertM(first, elements)
	first = deleteItem(first, 53)
	var actual [7]int
	var expected = [7]int{3, 6, 7, 2, 6, 3, 56}
	copy(actual[:], getElements(first))
	if actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReverseOneItem(t *testing.T) {
	var first *Node
	first = insertM(first, []int{3})
	first = reverse(first)
	var actual [1]int
	copy(actual[:], getElements(first))
	if actual != [1]int{3} {
		t.Errorf("Not reversed")
	}
}

func TestReverseTwoItem(t *testing.T) {
	var first *Node
	first = insertM(first, []int{3, 5})
	reversed := reverse(first)
	var actual [2]int
	copy(actual[:], getElements(reversed))
	expected := [2]int{5, 3}
	if actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReverseMulti(t *testing.T) {
	var first *Node
	first = insertM(first, []int{3, 5, 7, 2, 63, 32, 4})
	reversed := reverse(first)
	var actual [7]int
	copy(actual[:], getElements(reversed))
	expected := [7]int{4, 32, 63, 2, 7, 5, 3}
	if actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
