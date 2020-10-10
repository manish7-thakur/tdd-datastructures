package rabinkarp

import "testing"

func TestHashEmptyString(t *testing.T) {
	str := ""
	actual := hash(str)
	expected := 0
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashSingleChar(t *testing.T) {
	str := "a"
	actual := hash(str)
	expected := 97
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashDoubleChar(t *testing.T) {
	str := "ab"
	actual := hash(str)
	expected := 3105
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashTripleChar(t *testing.T) {
	str := "abc"
	actual := hash(str)
	expected := 96354
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashMultiChar(t *testing.T) {
	str := "~q7(@%s?2c"
	actual := hash(str)
	expected := -1289135674
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashLargeSting(t *testing.T) {
	str := "~q7(@%s?$hskKS63M>204KDA2k2c"
	actual := hash(str)
	expected := -1688811677
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashUsingPriorHashEmptyString(t *testing.T) {
	str := ""
	priorHash := hash(str)
	var next rune = 0
	actual := hash("")
	expected := newHash(priorHash, next, 0, 0)
	if actual != expected {
		t.Errorf("expected %d but found %d", actual, expected)
	}
}

func TestHashUsingPriorHashZeroSingleChar(t *testing.T) {
	str := ""
	priorHash := hash(str)
	next := 'a'
	actual := hash("a")
	expected := newHash(priorHash, next, 0, 1)
	if actual != expected {
		t.Errorf("expected %d but found %d", actual, expected)
	}
}

func TestHashUsingPriorHashNonZeroSingleChar(t *testing.T) {
	str := "b"
	priorHash := hash(str)
	next := 'a'
	actual := hash("a")
	expected := newHash(priorHash, next, 'b', 1)
	if actual != expected {
		t.Errorf("expected %d but found %d", actual, expected)
	}
}

func TestHashUsingPriorHashNonZeroDoubleChar(t *testing.T) {
	str := "bc"
	priorHash := hash(str)
	next := 'a'
	actual := hash("ca")
	expected := newHash(priorHash, next, 'b', 2)
	if actual != expected {
		t.Errorf("expected %d but found %d", actual, expected)
	}
}

func TestHashUsingPriorHashNonZeroTripleChar(t *testing.T) {
	str := "bcd"
	hashGen := New(str)
	next := 'a'
	actual := hash("cda")
	expected := hashGen.RollHash(next, 'b')
	if actual != expected {
		t.Errorf("expected %d but found %d", actual, expected)
	}
}
