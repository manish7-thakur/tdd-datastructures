package rabinkarp

import (
	"testing"
)

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
	expected := Hash("")
	hashGen.RollHash(next, 0)
	actual := hashGen.hash
	if expected != actual {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashUsingPriorHashZeroSingleChar(t *testing.T) {
	str := ""
	hashGen := New(str)
	next := 'a'
	expected := Hash("a")
	hashGen.RollHash(next, 0)
	actual := hashGen.hash
	if expected != actual {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashUsingPriorHashNonZeroSingleChar(t *testing.T) {
	str := "b"
	hashGen := New(str)
	next := 'a'
	expected := Hash("a")
	hashGen.RollHash(next, 'b')
	actual := hashGen.hash
	if expected != actual {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashUsingPriorHashNonZeroDoubleChar(t *testing.T) {
	str := "bc"
	hashGen := New(str)
	next := 'a'
	expected := Hash("ca")
	hashGen.RollHash(next, 'b')
	actual := hashGen.hash
	if expected != actual {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashUsingPriorHashNonZeroTripleChar(t *testing.T) {
	str := "bcd"
	hashGen := New(str)
	next := 'a'
	expected := Hash("cda")
	hashGen.RollHash(next, 'b')
	actual := hashGen.hash
	if expected != actual {
		t.Errorf("expected %d but found %d", actual, expected)
	}
}

func TestIndexRabinKarpEmptyString(t *testing.T) {
	actual := IndexRabinKarp("", "a")
	expected := -1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestIndexRabinKarpEmptySubString(t *testing.T) {
	actual := IndexRabinKarp("a", "")
	expected := 0
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestIndexRabinKarpSingleCharString(t *testing.T) {
	actual := IndexRabinKarp("a", "a")
	expected := 0
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestIndexRabinKarpDoubleCharString(t *testing.T) {
	actual := IndexRabinKarp("ab", "b")
	expected := 1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestIndexRabinKarpTripleCharString(t *testing.T) {
	actual := IndexRabinKarp("abc", "c")
	expected := 2
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestIndexRabinKarpFourCharStringDoubleCharSubstring(t *testing.T) {
	actual := IndexRabinKarp("abec", "be")
	expected := 1
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestIndexRabinKarpFiveCharStringDoubleCharSubstring(t *testing.T) {
	actual := IndexRabinKarp("dawec", "we")
	expected := 2
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestIndexRabinKarpMulticharStringMultiCharSubstring(t *testing.T) {
	actual := IndexRabinKarp("dawj389a/;,A$9AD2#903Md0-2@#JH2nx@xjhs02maOD#IXW30jwyKAsec.,:>:%#hjwilshwo#^&#@_+sjfEh", "#^&#@")
	expected := 74
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}
