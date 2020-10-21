package merge_sort

import (
	"testing"
)

func TestMergeEmptyArray(t *testing.T) {
	actual := merge([]int{}, []int{})
	if len(actual) != 0 {
		t.Errorf("expected empty but found %v", actual)
	}
}

func TestMergeEmptyAndOneItemArray(t *testing.T) {
	res := merge([]int{1}, []int{})
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{1}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeOneItemArrayAndEmpty(t *testing.T) {
	res := merge([]int{}, []int{2})
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeOneItemArrayEach(t *testing.T) {
	res := merge([]int{1}, []int{2})
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeOneItemArrayEachRandom(t *testing.T) {
	res := merge([]int{2}, []int{1})
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeTwoItemArrayEachRandom(t *testing.T) {
	res := merge([]int{2, 3}, []int{1, 4})
	actual := [4]int{}
	copy(actual[:], res)
	expected := [4]int{1, 2, 3, 4}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeDiffItemArrayEachRandom(t *testing.T) {
	res := merge([]int{2, 3, 5, 7}, []int{1, 4})
	actual := [6]int{}
	copy(actual[:], res)
	expected := [6]int{1, 2, 3, 4, 5, 7}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeSortEmptyArray(t *testing.T) {
	res := MergeSort([]int{})
	if len(res) != 0 {
		t.Errorf("expected empty but found %v", res)
	}
}

func TestMergeSortOneItemArray(t *testing.T) {
	res := MergeSort([]int{1})
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{1}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeSortTwoItemArray(t *testing.T) {
	res := MergeSort([]int{2, 1})
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeSortThreeItemArray(t *testing.T) {
	res := MergeSort([]int{2, 0, 1})
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{0, 1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeSortFourItemArray(t *testing.T) {
	res := MergeSort([]int{-1, 3, 0, 1})
	actual := [4]int{}
	copy(actual[:], res)
	expected := [4]int{-1, 0, 1, 3}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeSortMultiItemArray(t *testing.T) {
	res := MergeSort([]int{-33, -1, 321, -1, 23, 30, 1, 48, -263, 83, 0, 27, 26, 183, 171, 2362})
	actual := [16]int{}
	copy(actual[:], res)
	expected := [16]int{-263, -33, -1, -1, 0, 1, 23, 26, 27, 30, 48, 83, 171, 183, 321, 2362}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}
