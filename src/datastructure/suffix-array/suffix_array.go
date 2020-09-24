package suffix_array

import "sort"

func sortedSuffixArray(str string) []int {
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
