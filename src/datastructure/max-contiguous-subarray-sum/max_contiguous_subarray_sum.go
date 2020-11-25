package max_contiguous_subarray_sum

import (
	math2 "datastructure/math"
	"math"
)

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
	return math2.Max(math2.Max(max1, max2), max3)
}

func crossSum(arr []int, low, mid, high int) int {
	sum := 0
	leftSum := math.MinInt32
	rightSum := math.MinInt32
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
	return math2.Max(math2.Max(leftSum, rightSum), leftSum+rightSum)
}
