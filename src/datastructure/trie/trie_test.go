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
