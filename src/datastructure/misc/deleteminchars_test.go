package misc

import "testing"

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
}

func clearMap(occ map[int]int) {
	for k := range occ {
		delete(occ, k)
	}
}