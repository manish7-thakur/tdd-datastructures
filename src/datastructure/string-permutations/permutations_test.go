package string_permutations

import (
	"testing"
)

func TestPermutationsStringSingleChar(t *testing.T) {
	permutations := permutations("A")
	var actual [1]string
	copy(actual[:], permutations)
	if actual != [1]string{"A"} {
		t.Errorf("Expected empty but found %v", actual)
	}
}

func TestPermutationsStringTwoChar(t *testing.T) {
	permutations := permutations("AB")
	var actual [2]string
	copy(actual[:], permutations)
	expected := [2]string{"AB", "BA"}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestPermutationsStringThreeChar(t *testing.T) {
	permutations := permutations("ABC")
	var actual [6]string
	copy(actual[:], permutations)
	expected := [6]string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}