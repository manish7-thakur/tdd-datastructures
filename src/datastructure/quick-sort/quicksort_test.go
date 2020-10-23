package quick_sort

import "testing"

func TestPartitionEmptyArray(t *testing.T) {
	res := partition([]int{})
	if len(res) != 0 {
		t.Errorf("expected empty but found %v", res)
	}
}

func TestPartitionArrayOneElem(t *testing.T) {
	res := partition([]int{2})
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionArrayTwoElem(t *testing.T) {
	res := partition([]int{2, 1})
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func partition(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			temp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
		}
	}
	return arr
}
