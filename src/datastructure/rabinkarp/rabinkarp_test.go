package rabinkarp

import "testing"

func TestHashEmptyString(t *testing.T) {
	str := ""
	actual := Hash(str)
	expected := 0
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashSingleChar(t *testing.T) {
	str := "a"
	actual := Hash(str)
	expected := 97
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashDoubleChar(t *testing.T) {
	str := "ab"
	actual := Hash(str)
	expected := 3105
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashTripleChar(t *testing.T) {
	str := "abc"
	actual := Hash(str)
	expected := 96354
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashMultiChar(t *testing.T) {
	str := "~q7(@%s?2c"
	actual := Hash(str)
	expected := -1289135674
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashLargeSting(t *testing.T) {
	str := "~q7(@%s?$hskKS63M>204KDA2k2c"
	actual := Hash(str)
	expected := -1688811677
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashUsingPriorHashEmptyString(t *testing.T) {
	str := ""
	hashGen := New(str)
	var next rune = 0
	actual := Hash("")
	expected := hashGen.RollHash(next, 0)
	if actual != expected {
		t.Errorf("expected %d but found %d", actual, expected)
	}
}

func TestHashUsingPriorHashZeroSingleChar(t *testing.T) {
	str := ""
	hashGen := New(str)
	next := 'a'
	actual := Hash("a")
	expected := hashGen.RollHash(next, 0)
	if actual != expected {
		t.Errorf("expected %d but found %d", actual, expected)
	}
}

func TestHashUsingPriorHashNonZeroSingleChar(t *testing.T) {
	str := "b"
	hashGen := New(str)
	next := 'a'
	actual := Hash("a")
	expected := hashGen.RollHash(next, 'b')
	if actual != expected {
		t.Errorf("expected %d but found %d", actual, expected)
	}
}

func TestHashUsingPriorHashNonZeroDoubleChar(t *testing.T) {
	str := "bc"
	hashGen := New(str)
	next := 'a'
	actual := Hash("ca")
	expected := hashGen.RollHash(next, 'b')
	if actual != expected {
		t.Errorf("expected %d but found %d", actual, expected)
	}
}

func TestHashUsingPriorHashNonZeroTripleChar(t *testing.T) {
	str := "bcd"
	hashGen := New(str)
	next := 'a'
	actual := Hash("cda")
	expected := hashGen.RollHash(next, 'b')
	if actual != expected {
		t.Errorf("expected %d but found %d", actual, expected)
	}
}
