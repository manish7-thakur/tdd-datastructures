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
	res := make([]int, flen+slen)
	for i := 0; i < len(res); i++ {
		if j < flen {
			res[i] = first[j]
			j++
		} else if k < slen {
			res[i] = second[k]
			k++
		}
	}
	return res
}
