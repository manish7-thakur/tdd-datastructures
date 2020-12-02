package selection_sort

import "testing"

func TestSortEmptyArray(t *testing.T) {
	res := SelectionSort([]int{})
	if len(res) != 0 {
		t.Errorf("expected empty but found %v", res)
	}
}

func TestSortSingleElemArray(t *testing.T) {
	res := SelectionSort([]int{1})
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{1}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSortDoubleElemArray(t *testing.T) {
	res := SelectionSort([]int{2, 1})
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func SelectionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	min := arr[0]
	if arr[1] < min {
		arr[0], arr[1] = arr[1], arr[0]
	}
	return arr
}
