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
	res := CombSort([]int{3, 2, 1})
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{1, 2, 3}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, res)
	}
}

func TestSortingFourElemArray(t *testing.T) {
	res := CombSort([]int{4, 3, 2, 1})
	actual := [4]int{}
	copy(actual[:], res)
	expected := [4]int{1, 2, 3, 4}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, res)
	}
}

func TestSortingFiveElemArray(t *testing.T) {
	res := CombSort([]int{4, 3, 5, 2, 1})
	actual := [5]int{}
	copy(actual[:], res)
	expected := [5]int{1, 2, 3, 4, 5}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, res)
	}
}

func TestSortingMultiElemArray(t *testing.T) {
	res := CombSort([]int{47, 33, 56, 22, 11, -26, 83, -4, 55, -25, -216, 75, 24, 69, 8})
	actual := [15]int{}
	copy(actual[:], res)
	expected := [15]int{-216, -26, -25, -4, 8, 11, 22, 24, 33, 47, 55, 56, 69, 75, 83}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, res)
	}
}

func CombSort(arr []int) []int {
	arrlen := len(arr)
	width := arrlen / 2
	for i := 0; i < arrlen; i++ {
		for j := 0; j < arrlen-width; j++ {
			next := j + width
			if arr[j] > arr[next] {
				arr[j], arr[next] = arr[next], arr[j]
			}
		}
		width = width / 2
		if width == 0 {
			width = 1
		}
	}
	return arr
}
