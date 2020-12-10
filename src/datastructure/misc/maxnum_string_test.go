package misc

import (
	"testing"
)

func TestMaxTwoOrOneDigitNumPresent(t *testing.T) {
	actual := MaxNum("")
	expected := 0
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MaxNum("5")
	expected = 5
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MaxNum("05")
	expected = 5
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MaxNum("001")
	expected = 1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MaxNum("0010")
	expected = 10
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MaxNum("55")
	expected = 55
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MaxNum("505552")
	expected = 55
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
	actual = MaxNum("10101")
	expected = 10
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}
