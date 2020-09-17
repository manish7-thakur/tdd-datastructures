package lcs

import (
	"testing"
)

func TestEmptyStringLCSZero(t *testing.T) {
	actual := LCS("", "")
	expected := 0
	if actual != expected {
		t.Errorf("Expected %d but found %d", expected, actual)
	}
}

func TestSingleCharStringLCSZero(t *testing.T) {
	actual := LCS("A", "B")
	expected := 0
	if actual != expected {
		t.Errorf("Expected %d but found %d", expected, actual)
	}
}

func TestSingleCharStringLCSOne(t *testing.T) {
	actual := LCS("A", "A")
	expected := 1
	if actual != expected {
		t.Errorf("Expected %d but found %d", expected, actual)
	}
}

func TestDoubleCharStringLCSOne(t *testing.T) {
	actual := LCS("AB", "A")
	expected := 1
	if actual != expected {
		t.Errorf("Expected %d but found %d", expected, actual)
	}
}
func TestDoubleCharStringLCSTwo(t *testing.T) {
	actual := LCS("AB", "AB")
	expected := 2
	if actual != expected {
		t.Errorf("Expected %d but found %d", expected, actual)
	}
}

func TestMultiCharStringLCS3(t *testing.T) {
	actual := LCS("ABBABB", "BAABA")
	expected := 3
	if actual != expected {
		t.Errorf("Expected %d but found %d", expected, actual)
	}
}

func TestMultiCharStringLCS6(t *testing.T) {
	actual := LCS("U$NS(#NJDUS(@MS(@NA)#", "JDI%(S3SJE9#NS*#")
	expected := 6
	if actual != expected {
		t.Errorf("Expected %d but found %d", expected, actual)
	}
}
