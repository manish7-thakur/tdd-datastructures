package sos

import "testing"

func TestSosEmptyArray(t *testing.T) {
	actual := sos(1, []int{}, []int{})
	if len(actual) != 0 {
		t.Errorf("expected empty but found %v", actual)
	}
}

func TestSosZeroSum(t *testing.T) {
	actual := sos(0, []int{}, []int{})
	if len(actual) != 0 {
		t.Errorf("expected empty but found %v", actual)
	}
}

func TestSosSumLessThanAllElems(t *testing.T) {
	res := sos(1, []int{2}, []int{})
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{0}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}
func TestSosSumEqualToOneElems(t *testing.T) {
	res := sos(1, []int{1}, []int{})
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{1}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func sos(m int, n []int, curr []int) []int {
	setlen := len(n)
	if m == 0 || setlen == 0 {
		return curr
	} else if n[setlen-1] > m {
		return sos(m, n[0:setlen-1], curr)
	} else {
		curr = append(curr, n[setlen-1])
		sos(m-n[setlen-1], n[0:setlen-1], curr)
	}
	return curr
}
