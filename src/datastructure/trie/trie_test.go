package trie

import (
	"testing"
)

type TrieNode struct {
	chars     map[rune]TrieNode
	endOfWord bool
}

type Trie struct {
	first TrieNode
}

func (t *Trie) insert(s string) {
	if len(s) != 0 {
		char := rune(s[0])
		m := make(map[rune]TrieNode)
		m[char] = TrieNode{endOfWord: true}
		t.first = TrieNode{m, true}
	}
}

func TestInsertEmptyString(t *testing.T) {
	trie := Trie{}
	str := ""
	trie.insert(str)
	actual, ok := trie.first.chars['\000']
	if ok || actual.endOfWord {
		t.Errorf("expected empty but found %v", actual)
	}
}

func TestInsertSingleCharString(t *testing.T) {
	trie := Trie{}
	str := "a"
	trie.insert(str)
	actual, ok := trie.first.chars['a']
	if !ok || !actual.endOfWord {
		t.Errorf("expected %v but found %v", true, actual)
	}
}
