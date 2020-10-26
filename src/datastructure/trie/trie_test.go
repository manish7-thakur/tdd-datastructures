package trie

import (
	"testing"
)

func TestInsertEmptyString(t *testing.T) {
	trie := Trie{}
	str := ""
	trie.Insert(str)
	actual, ok := trie.first.chars['\000']
	if ok {
		t.Errorf("expected empty but found %v", actual)
	}
}

func TestInsertSingleCharString(t *testing.T) {
	trie := Trie{}
	str := "a"
	trie.Insert(str)
	actual, ok := trie.first.chars['a']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected true  but found %t", actual.endOfWord)
	}
}

func TestInsertDoubleCharString(t *testing.T) {
	trie := Trie{}
	str := "ab"
	trie.Insert(str)
	actual, ok := trie.first.chars['a']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != false {
		t.Errorf("expected false  but found %t", actual.endOfWord)
	}
	actual, ok = actual.chars['b']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected true  but found %t", actual.endOfWord)
	}
}

func TestInsertTripleCharString(t *testing.T) {
	trie := Trie{}
	str := "abc"
	trie.Insert(str)
	actual, ok := trie.first.chars['a']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != false {
		t.Errorf("expected false  but found %t", actual.endOfWord)
	}
	actual, ok = actual.chars['b']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != false {
		t.Errorf("expected false  but found %t", actual.endOfWord)
	}
	actual, ok = actual.chars['c']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected true  but found %t", actual.endOfWord)
	}
}

func TestInsertMultipleStrings(t *testing.T) {
	trie := Trie{}
	str1 := "a"
	str2 := "ab"
	str3 := "abc"
	trie.Insert(str1)
	trie.Insert(str2)
	trie.Insert(str3)
	actual, ok := trie.first.chars['a']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected true  but found %t", actual.endOfWord)
	}
	actual, ok = actual.chars['b']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected false  but found %t", actual.endOfWord)
	}
	actual, ok = actual.chars['c']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected false  but found %t", actual.endOfWord)
	}
}

func TestSearchEmptyString(t *testing.T) {
	str := ""
	trie := Trie{}
	actual := trie.Contains(str)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchSingleCharString(t *testing.T) {
	str := "a"
	trie := Trie{}
	trie.Insert(str)
	actual := trie.Contains(str)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharString(t *testing.T) {
	str := "ab"
	trie := Trie{}
	trie.Insert(str)
	actual := trie.Contains(str)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchSingleCharNonExistentString(t *testing.T) {
	str := "ab"
	trie := Trie{}
	trie.Insert(str)
	actual := trie.Contains("a")
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchTripleCharString(t *testing.T) {
	str := "aaaa"
	trie := Trie{}
	trie.Insert(str)
	actual := trie.Contains(str)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchMultiplePrefixedStrings(t *testing.T) {
	str1 := "a"
	trie := Trie{}
	trie.Insert(str1)
	actual := trie.Contains("aa")
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchMultipleNonExistentStrings(t *testing.T) {
	str1 := "a"
	str2 := "aaa"
	trie := Trie{}
	trie.Insert(str1)
	trie.Insert(str2)
	actual := trie.Contains(str1)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	actual = trie.Contains(str2)
	expected = true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	actual = trie.Contains("aa")
	expected = false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	actual = trie.Contains("aaaa")
	expected = false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}
