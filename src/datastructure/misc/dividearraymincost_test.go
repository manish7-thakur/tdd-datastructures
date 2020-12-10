package misc

import (
	"testing"
)

func TestMinCost(t *testing.T) {
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
