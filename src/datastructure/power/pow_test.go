package power

import "testing"

func TestPow0(t *testing.T) {
	res := pow(4, 0)
	expected := 1
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestPow1Of0(t *testing.T) {
	res := pow(0, 1)
	expected := 0
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestPow1Of1(t *testing.T) {
	res := pow(1, 1)
	expected := 1
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestPow1Of2(t *testing.T) {
	res := pow(2, 1)
	expected := 2
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestPow2Of2(t *testing.T) {
	res := pow(2, 2)
	expected := 4
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestPow4Of2(t *testing.T) {
	res := pow(2, 4)
	expected := 16
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestPow3Of2(t *testing.T) {
	res := pow(2, 3)
	expected := 8
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestPow5Of3(t *testing.T) {
	res := pow(3, 5)
	expected := 243
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestNearestPow2For3(t *testing.T) {
	res := nearestPow2(3)
	expected := 2
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestNearestPow2For5(t *testing.T) {
	res := nearestPow2(5)
	expected := 4
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestNearestPow2For6(t *testing.T) {
	res := nearestPow2(6)
	expected := 4
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestNearestPow2For7(t *testing.T) {
	res := nearestPow2(7)
	expected := 4
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestNearestPow2For8(t *testing.T) {
	res := nearestPow2(8)
	expected := 8
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestNearestPow2For9(t *testing.T) {
	res := nearestPow2(9)
	expected := 8
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestNearestPow2For15(t *testing.T) {
	res := nearestPow2(15)
	expected := 8
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestNearestPow2For17(t *testing.T) {
	res := nearestPow2(17)
	expected := 16
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}
