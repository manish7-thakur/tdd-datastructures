package max_contigous_subarray_sum

import (
	"datastructure/math"
	"testing"
)

func TestMaxContiguousSubarraySumEmpty(t *testing.T) {
	actual := MaxContiguousSubarraySum([]int{})
	expected := 0
	if actual != expected {
		t.Errorf("Expected %d but found %v", expected, actual)
	}
}

func TestMaxContiguousSubarraySumSingleElement(t *testing.T) {
	actual := MaxContiguousSubarraySum([]int{1})
	expected := 1
	if actual != expected {
		t.Errorf("Expected %d but found %v", expected, actual)
	}
}

func TestMaxContiguousSubarraySumSingleElementNegative(t *testing.T) {
	actual := MaxContiguousSubarraySum([]int{-1})
	expected := -1
	if actual != expected {
		t.Errorf("Expected %d but found %v", expected, actual)
	}
}

func TestMaxContiguousSubarraySumTwoElements(t *testing.T) {
	actual := MaxContiguousSubarraySum([]int{-1, 2})
	expected := 2
	if actual != expected {
		t.Errorf("Expected %d but found %v", expected, actual)
	}
}

func TestMaxContiguousSubarraySumThreeElements(t *testing.T) {
	actual := MaxContiguousSubarraySum([]int{1, 1, 2})
	expected := 4
	if actual != expected {
		t.Errorf("Expected %d but found %v", expected, actual)
	}
}

func MaxContiguousSubarraySum(arr []int) int {
	arrlen := len(arr)
	if arrlen == 0 {
		return 0
	}
	if arrlen == 1 {
		return arr[0]
	}
	mid := arrlen / 2
	max1 := MaxContiguousSubarraySum(arr[0:mid])
	max2 := MaxContiguousSubarraySum(arr[mid:arrlen])
	return math.Max(math.Max(max1, max2), max1+max2)
}
