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
		node, ok := next.chars[r]
		if !ok {
			//Insert into current node
			node = &TrieNode{}
			next.chars[r] = node
		}
		if i == len(s)-1 {
			node.endOfWord = true
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
	next := &t.first
	for _, r := range str {
		node, _ := next.chars[r]
		if len(node.chars) == 0 {
			delete(next.chars, r)
		}
	}
}
