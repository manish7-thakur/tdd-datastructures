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