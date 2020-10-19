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

func TestLPSDoubleCharDiffValue(t *testing.T) {
	res := LPS("ab")
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{0, 0}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestLPSTripleCharSameValue(t *testing.T) {
	res := LPS("aaa")
	actual := [3]int{}
	copy(actual[:], res)
	expected := [3]int{0, 1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestLPSFourCharDiffValue(t *testing.T) {
	res := LPS("aaba")
	actual := [4]int{}
	copy(actual[:], res)
	expected := [4]int{0, 1, 0, 1}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestLPSFiveCharDiffValue(t *testing.T) {
	res := LPS("abaab")
	actual := [5]int{}
	copy(actual[:], res)
	expected := [5]int{0, 0, 1, 1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}
func TestLPSSixCharDiffValue(t *testing.T) {
	res := LPS("abaaba")
	actual := [6]int{}
	copy(actual[:], res)
	expected := [6]int{0, 0, 1, 1, 2, 3}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestLPSMultiCharDiffValue(t *testing.T) {
	res := LPS("acacabacacabacacac")
	actual := [18]int{}
	copy(actual[:], res)
	expected := [18]int{0, 0, 1, 2, 3, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 4}
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
			j++
		} else {
			for ; j > 0 && str[i] != str[j]; {
				j = lpsarr[j-1]
			}
			if str[i] == str[j] {
				lpsarr[i] = j + 1
				j++
			}
		}
	}
	return lpsarr
}
