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

func TestPartitionIndexEmptyArray(t *testing.T) {
	idx, res := partitionIndex([]int{})
	expected := 0
	if idx != expected || len(res) != 0 {
		t.Errorf("expected %d but found %d", expected, idx)
	}
}

func TestPartitionIndexOneElem(t *testing.T) {
	idx, res := partitionIndex([]int{1})
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{1}
	if idx != 0 || expected != actual {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionIndexTwoElem(t *testing.T) {
	idx, res := partitionIndex([]int{2, 1})
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{1, 2}
	if idx != 1 || expected != actual {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionIndexTwoElemRepeated(t *testing.T) {
	idx, res := partitionIndex([]int{2, 1, 2})
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{1, 2, 2}
	if idx != 1 || expected != actual {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionIndexThreeElem(t *testing.T) {
	idx, res := partitionIndex([]int{3, 1, 2})
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{2, 1, 3}
	if idx != 2 || expected != actual {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionIndexThreeElemOrdered(t *testing.T) {
	idx, res := partitionIndex([]int{1, 2, 3})
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{1, 2, 3}
	if idx != 0 || expected != actual {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionIndexThreeElemOrderedReversed(t *testing.T) {
	idx, res := partitionIndex([]int{3, 2, 1})
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{1, 2, 3}
	if idx != 2 || expected != actual {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionIndexFourElem(t *testing.T) {
	idx, res := partitionIndex([]int{2, 4, 3, 1})
	actual := [4]int{}
	copy(actual[:], res)
	expected := [4]int{1, 2, 3, 4}
	if idx != 1 || expected != actual {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestPartitionIndexFourElemOrdered(t *testing.T) {
	idx, res := partitionIndex([]int{2, 4, 3, 1})
	actual := [4]int{}
	copy(actual[:], res)
	expected := [4]int{1, 2, 3, 4}
	if idx != 1 || expected != actual {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func partitionIndex(arr []int) (int, []int) {
	arrlen := len(arr)
	if arrlen < 2 {
		return 0, arr
	}
	pivot := arr[0]
	i := 0
	for j := 0; j < arrlen; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i], arr[0] = arr[0], arr[i]
	return i, arr
}
