package rabinkarp

import (
	"strings"
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
	actual := Hash("")
	expected := hashGen.RollHash(next, 0)
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashUsingPriorHashZeroSingleChar(t *testing.T) {
	str := ""
	hashGen := New(str)
	next := 'a'
	actual := Hash("a")
	expected := hashGen.RollHash(next, 0)
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashUsingPriorHashNonZeroSingleChar(t *testing.T) {
	str := "b"
	hashGen := New(str)
	next := 'a'
	actual := Hash("a")
	expected := hashGen.RollHash(next, 'b')
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashUsingPriorHashNonZeroDoubleChar(t *testing.T) {
	str := "bc"
	hashGen := New(str)
	next := 'a'
	actual := Hash("ca")
	expected := hashGen.RollHash(next, 'b')
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
	}
}

func TestHashUsingPriorHashNonZeroTripleChar(t *testing.T) {
	str := "bcd"
	hashGen := New(str)
	next := 'a'
	actual := Hash("cda")
	expected := hashGen.RollHash(next, 'b')
	if actual != expected {
		t.Errorf("expected %d but found %d", expected, actual)
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

func IndexRabinKarp(s string, substr string) int {
	substrlen := len(substr)
	slen := len(s)
	if substrlen == 0 {
		return 0
	} else if substrlen > slen {
		return -1
	} else {
		//compute hash substring
		//match with hash from string s part
		//if matches compare string else slide
		//if char matches return current idx
		hash := Hash(substr)
		hashGen := New(s[0:substrlen])
		if hashGen.hash == hash {
			if strings.Compare(substr, s[0:substrlen]) == 0 {
				return 0
			}
		}
	}
	return -1
}
