package misc

import (
	"math"
	"sort"
)

type Pair struct {
	Val int
	Idx int
}

func MinCostDivideArrayTwoParts(A []int) int {
	if len(A) < 5 {
		return 0
	}

	minarr := make([]*Pair, 3)
	minarr[0] = &Pair{A[1], 1}
	minarr[1] = &Pair{A[2], 2}
	minarr[2] = &Pair{A[3], 3}
	sort.Slice(minarr, func(i, j int) bool { return minarr[i].Val < minarr[j].Val })
	for i := 4; i < len(A)-1; i++ {
		if A[i] < minarr[2].Val {
			minarr[2].Val = A[i]
			minarr[2].Idx = i
			sort.Slice(minarr, func(i, j int) bool { return minarr[i].Val < minarr[j].Val })
		}
	}
	min1 := math.MaxInt32
	min2 := math.MaxInt32
	min3 := math.MaxInt32
	if Abs(minarr[0].Idx-minarr[1].Idx) != 1 {
		min1 = minarr[0].Val + minarr[1].Val
	}
	if Abs(minarr[0].Idx-minarr[2].Idx) != 1 {
		min2 = minarr[0].Val + minarr[2].Val
	}
	if Abs(minarr[1].Idx-minarr[2].Idx) != 1 {
		min3 = minarr[1].Val + minarr[2].Val
	}
	return Min(Min(min1, min2), min3)
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
