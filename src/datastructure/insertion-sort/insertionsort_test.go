package insertion_sort

import "testing"

func TestSortEmptyArray(t *testing.T) {
	res := InsertionSort([]int{})
	if len(res) != 0 {
		t.Errorf("expected empty but found %d", len(res))
	}
}

func TestSortSingleElemArray(t *testing.T) {
	res := InsertionSort([]int{1})
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{1}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSortDoubleElemArray(t *testing.T) {
	res := InsertionSort([]int{2, 1})
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSortTripleElemArray(t *testing.T) {
	res := InsertionSort([]int{3, 2, 1})
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{1, 2, 3}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSortFourElemArray(t *testing.T) {
	res := InsertionSort([]int{4, 3, 2, 1})
	actual := [4]int{}
	copy(actual[:], res)
	expected := [4]int{1, 2, 3, 4}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSortMultiElemArray(t *testing.T) {
	res := InsertionSort([]int{14, 23, 3, 47, 84, 67, 25, 59, 25, 347, 36})
	actual := [11]int{}
	copy(actual[:], res)
	expected := [11]int{3, 14, 23, 25, 25, 36, 47, 59, 67, 84, 347}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}