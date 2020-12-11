package misc

import (
	"testing"
)

func TestMinCostDivideArrayTwoParts(t *testing.T) {
	actual := MinCostDivideArrayTwoParts([]int{})
	expected := 0
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MinCostDivideArrayTwoParts([]int{5, 2, 4, 3, 3})
	expected = 5
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MinCostDivideArrayTwoParts([]int{5, 2, 1, 1, 3, 3})
	expected = 3
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MinCostDivideArrayTwoParts([]int{5, 1, 1, 2, 2, 3})
	expected = 3
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MinCostDivideArrayTwoParts([]int{5, 5, 2, 2, 1, 3})
	expected = 3
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MinCostDivideArrayTwoParts([]int{4, 10, 9, 8, 2, 8, 1, 3, 6, 3})
	expected = 3
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestMinNonAdjacentCost(t *testing.T) {
	actual := MinNonAdjacentCost([]*Pair{{1, 1}, &Pair{2, 2}, {4, 3}})
	expected := 5
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MinNonAdjacentCost([]*Pair{{1, 4}, {2, 5}, {4, 3}})
	expected = 6
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MinNonAdjacentCost([]*Pair{{1, 4}, {2, 2}, {3, 3}})
	expected = 3
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}

	actual = MinNonAdjacentCost([]*Pair{{2, 14}, {3, 15}, {6, 23}})
	expected = 8
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}
