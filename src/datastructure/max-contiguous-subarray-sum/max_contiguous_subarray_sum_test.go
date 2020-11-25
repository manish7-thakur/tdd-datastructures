package max_contiguous_subarray_sum

import (
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

func TestMaxContiguousSubarraySumThreeElementsNegative(t *testing.T) {
	actual := MaxContiguousSubarraySum([]int{2, -1, 2})
	expected := 3
	if actual != expected {
		t.Errorf("Expected %d but found %v", expected, actual)
	}
}

func TestMaxContiguousSubarraySumMultiElementsNegative(t *testing.T) {
	actual := MaxContiguousSubarraySum([]int{45, 23, -12, 2, -9, -1, 2, 42, 34})
	expected := 126
	if actual != expected {
		t.Errorf("Expected %d but found %v", expected, actual)
	}
}