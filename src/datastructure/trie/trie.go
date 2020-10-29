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
	for i, r := range s {
		if next.chars == nil {
			next.chars = make(map[rune]*TrieNode)
		}
		if node, ok := next.chars[r]; !ok {
			//Insert into current node
			node = &TrieNode{}
			if i == len(s)-1 {
				node.endOfWord = true
			}
			next.chars[r] = node
		} else {
			if i == len(s)-1 {
				node.endOfWord = true
			}
		}
		next = next.chars[r]
	}
}

func (t *Trie) Contains(str string, prefixSearch bool) bool {
	next := &t.first
	strlen := len(str)
	for i := 0; i < strlen; i++ {
		node, ok := next.chars[rune(str[i])]
		if !ok {
			return false
		} else if ok && i+1 == strlen {
			return prefixSearch || node.endOfWord
		}
		next = node
	}
	return false
}

func (t *Trie) Delete(str string) {
	if len(str) == 0 {
		return
	}
	delete(t.first.chars, rune(str[0]))
}
