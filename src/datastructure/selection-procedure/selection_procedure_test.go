package selection_procedure

import "testing"

func TestPartitionEmptyArray(t *testing.T) {
	actual := Partition([]int{})
	expected := -1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestPartitionSingleElem(t *testing.T) {
	actual := Partition([]int{1})
	expected := 0
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestPartitionDoubleElem(t *testing.T) {
	actual := Partition([]int{2, 1})
	expected := 1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func Partition(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	pos := 0
	pivot := arr[0]
	for i, n := range arr {
		if n < pivot {
			arr[pos], arr[i] = arr[i], arr[pos]
			pos++
		}
	}
	return pos
}
