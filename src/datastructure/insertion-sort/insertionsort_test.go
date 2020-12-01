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

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		if arr[i] < arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}
