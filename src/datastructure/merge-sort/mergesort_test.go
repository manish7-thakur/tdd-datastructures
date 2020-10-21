package merge_sort

import (
	"testing"
)

func TestMergeEmptyArray(t *testing.T) {
	actual := merge([]int{}, []int{})
	if len(actual) != 0 {
		t.Errorf("expected empty but found %v", actual)
	}
}

func TestMergeEmptyAndOneItemArray(t *testing.T) {
	res := merge([]int{1}, []int{})
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{1}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeOneItemArrayAndEmpty(t *testing.T) {
	res := merge([]int{}, []int{2})
	actual := [1]int{}
	copy(actual[:], res)
	expected := [1]int{2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeOneItemArrayEach(t *testing.T) {
	res := merge([]int{1}, []int{2})
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeOneItemArrayEachRandom(t *testing.T) {
	res := merge([]int{2}, []int{1})
	actual := [2]int{}
	copy(actual[:], res)
	expected := [2]int{1, 2}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeTwoItemArrayEachRandom(t *testing.T) {
	res := merge([]int{2, 3}, []int{1, 4})
	actual := [4]int{}
	copy(actual[:], res)
	expected := [4]int{1, 2, 3, 4}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func TestMergeDiffItemArrayEachRandom(t *testing.T) {
	res := merge([]int{2, 3, 5, 7}, []int{1, 4})
	actual := [6]int{}
	copy(actual[:], res)
	expected := [6]int{1, 2, 3, 4, 5, 7}
	if actual != expected {
		t.Errorf("expected %v but found %v", expected, actual)
	}
}

func merge(first []int, second []int) []int {
	flen := len(first)
	if flen == 0 {
		return second
	}
	slen := len(second)
	if slen == 0 {
		return first
	}
	j := 0
	k := 0
	reslen := flen + slen
	res := make([]int, reslen)
	for i := 0; i < reslen; i++ {
		if j < flen && k < slen {
			if first[j] < second[k] {
				res[i] = first[j]
				j++
			} else {
				res[i] = second[k]
				k++
			}
		} else if j < flen {
			res[i] = first[j]
			j++
		} else {
			res[i] = second[k]
			k++
		}
	}
	return res
}
