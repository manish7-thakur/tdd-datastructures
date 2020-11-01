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
	strlen := len(str)
	if strlen == 0 {
		return
	}
	stack := make([]*TrieNode, 0, strlen)
	charStack := make([]rune, 0, strlen)
	curr := &t.first
	for i, r := range str {
		node, _ := curr.chars[r]
		if i+1 == strlen {
			node.endOfWord = false
		}
		stack = append(stack, curr)
		charStack = append(charStack, r)
		curr = node
	}
	for ; len(stack) != 0; {
		node := stack[len(stack)-1]
		r := charStack[len(stack)-1]

		if !node.chars[r].endOfWord {
			delete(node.chars, r)
		}
		if len(node.chars) == 0 {
			node = nil
		}
		stack = stack[:len(stack)-1]
	}
}
