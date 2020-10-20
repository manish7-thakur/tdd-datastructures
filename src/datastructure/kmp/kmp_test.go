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
	res := LPS("bbabbb")
	actual := [6]int{}
	copy(actual[:], res)
	expected := [6]int{0, 1, 0, 1, 2, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestLPSSevenCharDiffValue(t *testing.T) {
	res := LPS("bbbabb")
	actual := [6]int{}
	copy(actual[:], res)
	expected := [6]int{0, 1, 2, 0, 1, 2}
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

func TestKMPSearchEmptyString(t *testing.T) {
	actual := KMPSearch("", "")
	expected := 0
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestKMPSearchSingleCharString(t *testing.T) {
	actual := KMPSearch("a", "b")
	expected := -1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func KMPSearch(text string, pattern string) int {
	patlen := len(pattern)
	if patlen == 0 {
		return 0
	}
	j := 0
	textlen := len(text)
	for i := 0; i < textlen; i++ {
		if pattern[j] == text[i] {
			return i
		}
	}
	return -1
}
