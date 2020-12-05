package comb_sort

import "testing"

func TestSortingEmptyArray(t *testing.T) {
	res := CombSort([]int{})
	if len(res) != 0 {
		t.Errorf("expected empty but found %v", res)
	}
}

func TestSortingSingleElemArray(t *testing.T) {
	res := CombSort([]int{1})
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{1}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, res)
	}
}

func TestSortingDoubleElemArray(t *testing.T) {
	res := CombSort([]int{2, 1})
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, res)
	}
}

func TestSortingTripleElemArray(t *testing.T) {
	res := CombSort([]int{2, 1})
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, res)
	}
}

func CombSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	if arr[0] > arr[1] {
		arr[0], arr[1] = arr[1], arr[0]
	}
	return arr
}
