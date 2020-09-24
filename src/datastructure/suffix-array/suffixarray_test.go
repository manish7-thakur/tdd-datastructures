package suffix_array

import (
	"sort"
	"testing"
)

func TestSuffixArrayEmptyString(t *testing.T) {
	str := ""
	sa := sortedSuffixArray(str)
	actual := [0]int{}
	copy(actual[:], sa)
	if actual != [0]int{} {
		t.Errorf("expected empty but found %v", actual)
	}
}

func TestSuffixArraySingleChar(t *testing.T) {
	str := "e"
	sa := sortedSuffixArray(str)
	actual := [1]int{}
	copy(actual[:], sa)
	expected := [1]int{0}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSuffixArrayTwoChar(t *testing.T) {
	str := "ea"
	sa := sortedSuffixArray(str)
	actual := [2]int{}
	copy(actual[:], sa)
	expected := [2]int{1, 0}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSuffixArrayThreeChar(t *testing.T) {
	str := "bea"
	sa := sortedSuffixArray(str)
	actual := [3]int{}
	copy(actual[:], sa)
	expected := [3]int{2, 0, 1}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

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
