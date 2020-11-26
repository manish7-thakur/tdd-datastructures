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

func TestPartitionDoubleElems(t *testing.T) {
	actual := Partition([]int{2, 1})
	expected := 1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestPartitionTripleElems(t *testing.T) {
	actual := Partition([]int{3, 2, 1})
	expected := 2
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestPartitionFourElems(t *testing.T) {
	actual := Partition([]int{3, 1, 4, 2})
	expected := 2
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestPartitionFiveElems(t *testing.T) {
	actual := Partition([]int{4, 5, 1, 2, 3})
	expected := 3
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestSelectionKthElemEmptyArray(t *testing.T) {
	actual := SelectElem([]int{}, 2)
	expected := -1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestSelectionKthElemSingleElemArray(t *testing.T) {
	actual := SelectElem([]int{2}, 1)
	expected := 2
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestSelectionKthElemDoubleElemArray(t *testing.T) {
	actual := SelectElem([]int{4, 3}, 1)
	expected := 3
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestSelectionKthElemTripleElemArray(t *testing.T) {
	actual := SelectElem([]int{5, 2, 4}, 1)
	expected := 2
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestSelectionKthElemFourElemArray(t *testing.T) {
	actual := SelectElem([]int{5, 3, 2, 4}, 2)
	expected := 3
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func SelectElem(arr []int, k int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 && k == 1 {
		return arr[0]
	}
	m := Partition(arr)
	for ; m != 0 && m != k; {
		if m == k {
			return arr[m]
		} else if k > m {
			m = Partition(arr[m+1:])
		} else {
			m = Partition(arr[:m])
		}
	}
	return arr[m-1]
}
