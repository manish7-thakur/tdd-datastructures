package power

import "testing"

func TestPowZero(t *testing.T) {
	res := pow(4, 0)
	expected := 1
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestPowOneOfZero(t *testing.T) {
	res := pow(0, 1)
	expected := 0
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestPowOneOfOne(t *testing.T) {
	res := pow(1, 1)
	expected := 1
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestPowOneOfTwo(t *testing.T) {
	res := pow(2, 1)
	expected := 2
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestPowTwoOfTwo(t *testing.T) {
	res := pow(2, 2)
	expected := 4
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestPowTwoOfFour(t *testing.T) {
	res := pow(2, 4)
	expected := 16
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}
