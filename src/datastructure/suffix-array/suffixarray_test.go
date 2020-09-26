package suffix_array

import (
	"testing"
)

func TestSuffixArrayEmptyString(t *testing.T) {
	str := ""
	sa := New(str)
	actual := [0]int{}
	copy(actual[:], sa.suffixSorted)
	if actual != [0]int{} {
		t.Errorf("expected empty but found %v", actual)
	}
}

func TestSuffixArraySingleChar(t *testing.T) {
	str := "e"
	sa := New(str)
	actual := [1]int{}
	copy(actual[:], sa.suffixSorted)
	expected := [1]int{0}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSuffixArrayTwoChar(t *testing.T) {
	str := "ea"
	sa := New(str)
	actual := [2]int{}
	copy(actual[:], sa.suffixSorted)
	expected := [2]int{1, 0}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSuffixArrayThreeChar(t *testing.T) {
	str := "bea"
	sa := New(str)
	actual := [3]int{}
	copy(actual[:], sa.suffixSorted)
	expected := [3]int{2, 0, 1}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSuffixArrayThreeCharIndx(t *testing.T) {
	str := "bea"
	sa := New(str)
	actual := [30]int{}
	copy(actual[:], sa.suffixSorted)
	expected := [30]int{2, 0, 1}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}
