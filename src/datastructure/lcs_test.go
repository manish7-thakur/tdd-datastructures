package datastructure

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

func LCS(s1 string, s2 string) int {
	return lcsCount(s1, s2, len(s1), len(s2))
}

func lcsCount(s1 string, s2 string, i int, j int) int {
	if i < 1 || j < 1 {
		return 0
	} else if s1[i-1] == s2[j-1] {
		return 1 + lcsCount(s1, s2, i-1, j-1)
	} else {
		return Max(lcsCount(s1, s2, i-1, j), lcsCount(s1, s2, i, j-1))
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
