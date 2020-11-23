package max_contigous_subarray_sum

import (
	"datastructure/math"
	math2 "math"
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

func MaxContiguousSubarraySum(arr []int) int {
	arrlen := len(arr)
	if arrlen == 0 {
		return 0
	}
	if arrlen == 1 {
		return arr[0]
	}
	mid := arrlen / 2
	max1 := MaxContiguousSubarraySum(arr[:mid])
	max2 := MaxContiguousSubarraySum(arr[mid+1:])
	max3 := crossSum(arr, 0, mid, arrlen-1)
	return math.Max(math.Max(max1, max2), max3)
}

func crossSum(arr []int, low, mid, high int) int {
	sum := 0
	leftSum := math2.MinInt32
	rightSum := math2.MinInt32
	for i := mid; i >= low; i-- {
		sum += arr[i]
		if sum > leftSum {
			leftSum = sum
		}
	}
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum += arr[i]
		if sum > rightSum {
			rightSum = sum
		}
	}
	return math.Max(math.Max(leftSum, rightSum), leftSum+rightSum)
}
