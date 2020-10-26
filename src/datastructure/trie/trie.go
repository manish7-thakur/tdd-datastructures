package trie

type TrieNode struct {
	chars     map[rune]*TrieNode
	endOfWord bool
}

type Trie struct {
	first TrieNode
}

func (t *Trie) Insert(s string) {
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
			//Insert into current node
			node := TrieNode{endOfWord: endOfWord}
			next.chars[r] = &node
		}
		next = next.chars[r]
	}
}

func (t *Trie) Contains(str string) bool {
	next := &t.first
	for i := 0; i < len(str); i++ {
		node, ok := next.chars[rune(str[i])]
		if ok && i+1 == len(str) && node.endOfWord {
			return true
		}
		next = node
	}
	return false
}
