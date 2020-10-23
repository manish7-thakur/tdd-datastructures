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

func TestPartitionArrayTwoElemOrdered(t *testing.T) {
	res := partition([]int{1, 2})
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionArrayThreeElem(t *testing.T) {
	res := partition([]int{2, 3, 1})
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{1, 2, 3}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionArrayThreeElemOrdered(t *testing.T) {
	res := partition([]int{1, 2, 3})
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{1, 2, 3}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionArrayThreeElemTwoEqual(t *testing.T) {
	res := partition([]int{2, 1, 2})
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{1, 2, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionArrayFourElem(t *testing.T) {
	res := partition([]int{2, 4, 3, 1})
	actual := [4]int{}
	copy(actual[:], res)
	expected := [4]int{1, 2, 3, 4}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionArrayFourElemOrdered(t *testing.T) {
	res := partition([]int{1, 2, 3, 4})
	actual := [4]int{}
	copy(actual[:], res)
	expected := [4]int{1, 2, 3, 4}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionArrayFiveElem(t *testing.T) {
	res := partition([]int{3, 1, 2, 5, 4})
	actual := [5]int{}
	copy(actual[:], res)
	expected := [5]int{2, 1, 3, 5, 4}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionArrayFiveElemOrdered(t *testing.T) {
	res := partition([]int{5, 4, 3, 2, 1})
	actual := [5]int{}
	copy(actual[:], res)
	expected := [5]int{1, 4, 3, 2, 5}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}
