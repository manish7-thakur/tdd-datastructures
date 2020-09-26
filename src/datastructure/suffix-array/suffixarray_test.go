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

func TestSuffixArrayMultiChar(t *testing.T) {
	str := "erywuemnbzcsadhihwueiuqwyeiqwyezskzjdksjhfsdjkfhsdgfhsgdfjhsdgfjsgfshdfgsjhfgsdwiuewfxaLKDJA"
	s1 := New(str)
	s2 := SortedSuffixArray(str)
	s1res := [92]int{}
	s2res := [92]int{}
	copy(s1res[:], s1.suffixSorted)
	copy(s2res[:], s2)
	if s1res != s2res {
		t.Errorf("expected %v but found %v", s1res, s2res)
	}
}

func BenchmarkIndex_New(b *testing.B) {
	str := "erywuemnbzcsadhihwueiuqwyeiqwyezskzjdksjhfsdjkfhsdgfhsgdfjhsdgfjsgfshdfgsjhfgsdwiuewfxaLKDJA"
	for i := 0; i < b.N; i++ {
		New(str)
	}
}

func BenchmarkIndex_SortedSuffixArray(b *testing.B) {
	str := "erywuemnbzcsadhihwueiuqwyeiqwyezskzjdksjhfsdjkfhsdgfhsgdfjhsdgfjsgfshdfgsjhfgsdwiuewfxaLKDJA"
	for i := 0; i < b.N; i++ {
		SortedSuffixArray(str)
	}
}
