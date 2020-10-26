package trie

import (
	"testing"
)

type TrieNode struct {
	chars     map[rune]*TrieNode
	endOfWord bool
}

type Trie struct {
	first TrieNode
}

func (t *Trie) insert(s string) {
	var next = &t.first
	endOfWord := false
	for i, r := range s {
		if next.chars == nil {
			next.chars = make(map[rune]*TrieNode)
		}
		if i == len(s)-1 {
			endOfWord = true
		}
		if _, ok := next.chars[r]; !ok {
			//insert into current node
			node := TrieNode{endOfWord: endOfWord}
			next.chars[r] = &node
		}
		next = next.chars[r]
	}
}

func TestInsertEmptyString(t *testing.T) {
	trie := Trie{}
	str := ""
	trie.insert(str)
	actual, ok := trie.first.chars['\000']
	if ok {
		t.Errorf("expected empty but found %v", actual)
	}
}

func TestInsertSingleCharString(t *testing.T) {
	trie := Trie{}
	str := "a"
	trie.insert(str)
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
	trie.insert(str)
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
	trie.insert(str)
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
	trie.insert(str1)
	trie.insert(str2)
	trie.insert(str3)
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
