package misc

import (
	"testing"
)

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

func TestOccurrenceMap(t *testing.T) {
	occMap := FindOccurrence("a")
	if occMap[1] != 1 || len(occMap) != 1 {
		t.Errorf("mismatch found")
	}
	occMap = FindOccurrence("aa")
	if occMap[2] != 1 || len(occMap) != 1 {
		t.Errorf("mismatch found")
	}
	occMap = FindOccurrence("aab")
	if occMap[1] != 1 || occMap[2] != 1 || len(occMap) != 2 {
		t.Errorf("mismatch found")
	}
	occMap = FindOccurrence("aabb")
	if occMap[2] != 2 || len(occMap) != 1 {
		t.Errorf("mismatch found")
	}
	occMap = FindOccurrence("aabbbcccc")
	if occMap[2] != 1 || occMap[3] != 1 || occMap[4] != 1 || len(occMap) != 3 {
		t.Errorf("mismatch found")
	}
	occMap = FindOccurrence("ccaaffddecee")
	if occMap[2] != 3 || occMap[3] != 2 || len(occMap) != 2 {
		t.Errorf("mismatch found")
	}
	occMap = FindOccurrence("eeee")
	if occMap[4] != 1 || len(occMap) != 1 {
		t.Errorf("mismatch found")
	}
	occMap = FindOccurrence("example")
	if occMap[1] != 5 || occMap[2] != 1 || len(occMap) != 2 {
		t.Errorf("mismatch found")
	}
}

func TestMinCharsToDelete(t *testing.T) {
	actual := MinCharsToDelete("ccaaffddecee")
	expected := 6
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MinCharsToDelete("aaaabbbb")
	expected = 1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MinCharsToDelete("eeee")
	expected = 0
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MinCharsToDelete("example")
	expected = 4
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func clearMap(occ map[int]int) {
	for k := range occ {
		delete(occ, k)
	}
}
