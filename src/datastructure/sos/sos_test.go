package sos

import "testing"

func TestSosEmptyArray(t *testing.T) {
	actual := sos(1, []int{})
	expected := false
	if expected != actual {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSosZeroSum(t *testing.T) {
	actual := sos(0, []int{})
	expected := true
	if expected != actual {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSosOneElem(t *testing.T) {
	actual := sos(1, []int{1})
	expected := true
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSosTwoElem(t *testing.T) {
	actual := sos(3, []int{1, 2})
	expected := true
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func sos(m int, n []int) bool {
	if m == 0 {
		return true
	}
	if len(n) == 0 {
		return false
	}
	return sos(m-n[0], n[1:]) || sos(m, n[1:])
}
