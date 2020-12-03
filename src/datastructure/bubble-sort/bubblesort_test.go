package bubble_sort

import "testing"

func TestSortEmptyArray(t *testing.T) {
	res := BubbleSort([]int{})
	if len(res) != 0 {
		t.Errorf("expected empty but found %v", res)
	}
}

func TestSortSingleElemArray(t *testing.T) {
	res := BubbleSort([]int{1})
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{1}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSortDoubleElemArray(t *testing.T) {
	res := BubbleSort([]int{2, 1})
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSortTripleElemArray(t *testing.T) {
	res := BubbleSort([]int{2, 3, 1})
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{1, 2, 3}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSortFourElemArray(t *testing.T) {
	res := BubbleSort([]int{2, 4, 3, 1})
	actual := [4]int{}
	copy(actual[:], res)
	expected := [4]int{1, 2, 3, 4}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSortMultiElemArray(t *testing.T) {
	res := BubbleSort([]int{24, 14, 53, 31, 74, 24, 46, 7, 235, 57, 34, 56, 26})
	actual := [13]int{}
	copy(actual[:], res)
	expected := [13]int{7, 14, 24, 24, 26, 31, 34, 46, 53, 56, 57, 74, 235}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}