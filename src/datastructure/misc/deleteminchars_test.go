package misc

import (
	"testing"
)

func TestDeleteMinCharsArray(t *testing.T) {
	actual := DeleteMinChars([]int{2})
	expected := 0
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = DeleteMinChars([]int{2, 2})
	expected = 1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = DeleteMinChars([]int{2, 2, 2})
	expected = 3
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = DeleteMinChars([]int{2, 2, 2, 2})
	expected = 5
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = DeleteMinChars([]int{2, 3})
	expected = 0
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = DeleteMinChars([]int{2, 2, 3})
	expected = 1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = DeleteMinChars([]int{2, 2, 3, 3})
	expected = 4
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = DeleteMinChars([]int{2, 2, 2, 3, 3})
	expected = 6
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = DeleteMinChars([]int{2, 2, 3, 3, 3})
	expected = 7
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = DeleteMinChars([]int{2, 2, 3, 3, 3, 4})
	expected = 7
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}

	actual = DeleteMinChars([]int{2, 2, 3, 3, 3, 4, 4})
	expected = 11
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = DeleteMinChars([]int{1, 1, 1, 1, 1, 2})
	expected = 4
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = DeleteMinChars([]int{1, 1, 1, 1, 1, 3, 3})
	expected = 5
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = DeleteMinChars([]int{1, 1, 3, 3, 5, 5})
	expected = 3
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestDeleteMinCharsMap(t *testing.T) {

	occ := make(map[int]int)
	occ[2] = 1
	actual := DeleteMinCharsMap(occ)
	expected := 0
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	clearMap(occ)
	occ[2] = 2
	actual = DeleteMinCharsMap(occ)
	expected = 1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	clearMap(occ)
	occ[2] = 3
	actual = DeleteMinCharsMap(occ)
	expected = 3
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	clearMap(occ)
	occ[2] = 4
	actual = DeleteMinCharsMap(occ)
	expected = 5
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	clearMap(occ)
	occ[2] = 1
	occ[3] = 1
	actual = DeleteMinCharsMap(occ)
	expected = 0
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	clearMap(occ)
	occ[2] = 2
	occ[3] = 1
	actual = DeleteMinCharsMap(occ)
	expected = 1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	clearMap(occ)
	occ[2] = 2
	occ[3] = 2
	actual = DeleteMinCharsMap(occ)
	expected = 4
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	clearMap(occ)
	occ[2] = 3
	occ[3] = 2
	actual = DeleteMinCharsMap(occ)
	expected = 6
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	clearMap(occ)
	occ[2] = 2
	occ[3] = 3
	actual = DeleteMinCharsMap(occ)
	expected = 7
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	clearMap(occ)
	occ[2] = 2
	occ[3] = 3
	occ[4] = 1
	actual = DeleteMinCharsMap(occ)
	expected = 7
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	clearMap(occ)
	occ[2] = 2
	occ[3] = 3
	occ[4] = 2
	actual = DeleteMinCharsMap(occ)
	expected = 11
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	clearMap(occ)
	occ[1] = 5
	occ[2] = 1
	actual = DeleteMinCharsMap(occ)
	expected = 4
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	clearMap(occ)
	occ[1] = 5
	occ[3] = 2
	actual = DeleteMinCharsMap(occ)
	expected = 5
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	clearMap(occ)
	occ[1] = 2
	occ[3] = 2
	occ[5] = 2
	actual = DeleteMinCharsMap(occ)
	expected = 3
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestOccuranceArray(t *testing.T) {
	occMap := FindOccurrence("a")
	expected := 1
	actual := occMap[1]
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	occMap = FindOccurrence("aa")
	expected = 1
	actual = occMap[2]
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	_, ok := occMap[1]
	if ok {
		t.Errorf("expected to be deleted but was present")
	}
	occMap = FindOccurrence("aab")
	if occMap[1] != 1 && occMap[2] != 1 && len(occMap) != 2 {
		t.Errorf("mismatch found")
	}
	occMap = FindOccurrence("aabb")
	if occMap[2] != 2 && len(occMap) != 1 {
		t.Errorf("mismatch found")
	}
	occMap = FindOccurrence("aabbbcccc")
	if occMap[2] != 1 && occMap[3] != 1 && occMap[4] != 1 && len(occMap) != 3 {
		t.Errorf("mismatch found")
	}
	occMap = FindOccurrence("ccaaffddecce")
	if occMap[2] != 4 && occMap[4] != 1 && len(occMap) != 2 {
		t.Errorf("mismatch found")
	}
	occMap = FindOccurrence("eeee")
	if occMap[4] != 1 && len(occMap) != 2 {
		t.Errorf("mismatch found")
	}
	occMap = FindOccurrence("example")
	if occMap[1] != 5 && occMap[2] != 1 && len(occMap) != 2 {
		t.Errorf("mismatch found")
	}
}

func FindOccurrence(s string) map[int]int {
	charMap := make(map[rune]int)
	occMap := make(map[int]int)
	for _, c := range s {
		if oldVal, ok := charMap[c]; ok {
			if oldVal == 1 {
				delete(occMap, oldVal)
			} else {
				occMap[oldVal]--
			}
		}
		charMap[c]++
		occMap[charMap[c]]++
	}
	return occMap
}

func clearMap(occ map[int]int) {
	for k := range occ {
		delete(occ, k)
	}
}
