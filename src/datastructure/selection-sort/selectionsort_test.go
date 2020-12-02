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

func TestSortTripleElemArray(t *testing.T) {
	res := SelectionSort([]int{2, 3, 1})
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{1, 2, 3}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSortFourElemArray(t *testing.T) {
	res := SelectionSort([]int{2, 4, 3, 1})
	actual := [4]int{}
	copy(actual[:], res)
	expected := [4]int{1, 2, 3, 4}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSortMultipleElemArray(t *testing.T) {
	res := SelectionSort([]int{24, 42, 73, 21, 37, 66, 10, 35, 46, 29, 45, 45, 23, 69, 34})
	actual := [15]int{}
	copy(actual[:], res)
	expected := [15]int{1, 2, 3, 4}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func SelectionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		minidx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minidx] {
				minidx = j
			}
		}
		arr[i], arr[minidx] = arr[minidx], arr[i]
	}

	return arr
}
