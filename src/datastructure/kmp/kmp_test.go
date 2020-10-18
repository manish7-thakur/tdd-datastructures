package kmp

import "testing"

func TestLPSEmptyString(t *testing.T) {
	actual := LPS("")
	if len(actual) != 0 {
		t.Errorf("expected empty but found %v", actual)
	}
}

func TestLPSSingleChar(t *testing.T) {
	res := LPS("a")
	actual := [1]int{1}
	copy(actual[:], res)
	expected := [1]int{0}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestLPSDoubleCharSameValue(t *testing.T) {
	res := LPS("aa")
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{0, 1}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func LPS(str string) []int {
	strlen := len(str)
	if strlen == 0 {
		return []int{}
	}
	lpsarr := make([]int, strlen)
	j := 0
	for i := 1; i < strlen; i++ {
		if str[i] == str[j] {
			lpsarr[i] = j + 1
		}
	}
	return lpsarr
}
