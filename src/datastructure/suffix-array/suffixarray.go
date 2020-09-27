package suffix_array

import (
	"sort"
)

func SortedSuffixArray(str string) []int {
	stringLen := len(str)
	suffixArr := make([]int, stringLen)
	suffixArrStr := make([]string, stringLen)
	suffixStartIdx := make(map[string]int, stringLen)
	for i := 0; i < stringLen; i++ {
		suffix := str[i:]
		suffixStartIdx[suffix] = i
		suffixArrStr[i] = suffix
	}
	sort.Strings(suffixArrStr)
	for i, suffix := range suffixArrStr {
		suffixArr[i] = suffixStartIdx[suffix]
	}
	return suffixArr
}

type Index struct {
	str          string
	suffixSorted []int
	lcp          []int
}

func New(str string) Index {
	suffixSorted := make([]int, len(str))
	lcp := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		suffixSorted[i] = i
	}
	index := Index{str, suffixSorted, lcp}
	sort.Sort(index)
	for i := 1; i < len(lcp); i++ {
		lcp[i] = 0
	}
	return index
}

func (idx Index) Len() int {
	return len(idx.suffixSorted)
}

func (idx Index) Less(i, j int) bool {
	str1 := idx.str[idx.suffixSorted[i]:]
	str2 := idx.str[idx.suffixSorted[j]:]
	return str1 < str2
}

func (idx Index) Swap(i, j int) {
	idx.suffixSorted[i], idx.suffixSorted[j] = idx.suffixSorted[j], idx.suffixSorted[i]
}

func LCP(str1 string, str2 string) int {
	lcpLenth := 0
	strlen := 0
	if len(str1) < len(str2) {
		strlen = len(str1)
	} else {
		strlen = len(str2)
	}
	for i := 0; i < strlen; i++ {
		if str1[i] == str2[i] {
			lcpLenth++
		} else {
			return lcpLenth
		}
	}
	return lcpLenth
}

