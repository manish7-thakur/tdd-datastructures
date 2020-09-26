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
}

func New(str string) Index {
	suffixSorted := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		suffixSorted[i] = i
	}
	index := Index{str, suffixSorted}
	sort.Sort(index)
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
