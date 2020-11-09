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

func TestSosOneElemSumPresent(t *testing.T) {
	actual := sos(1, []int{1})
	expected := true
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSosOneElemSumNotPresent(t *testing.T) {
	actual := sos(1, []int{2})
	expected := false
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSosTwoElemSumPresent(t *testing.T) {
	actual := sos(3, []int{1, 2})
	expected := true
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSosTwoElemSumNotPresent(t *testing.T) {
	actual := sos(3, []int{1, 1})
	expected := false
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSosMultiElemSumPresent(t *testing.T) {
	actual := sos(87, []int{52, 41, 46, 9, 3, 12, 83, 32, 42, 85, 31, 6, 3, 55})
	expected := true
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestSosMultiElemNotPresent(t *testing.T) {
	actual := sos(10, []int{52, 41, 46, 9, 3, 12, 83, 32, 42, 85, 31, 6, 3, 55})
	expected := false
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
