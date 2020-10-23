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

func TestPartitionArrayThreeElem(t *testing.T) {
	res := partition([]int{2, 3, 1})
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{1, 2, 3}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func partition(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	j := len(arr) - 1
	i := 1
	pivot := arr[0]
	for ; i < j; {
		for ; i < len(arr) && arr[i] < pivot; {
			i++
		}
		for ; j > 0 && arr[j] >= pivot; {
			j--
		}
		swap(&arr[i], &arr[j])
		i++
		j--
	}
	swap(&arr[0], &arr[j])

	return arr
}

func swap(i *int, j *int) {
	temp := *i
	*i = *j
	*j = temp
}
