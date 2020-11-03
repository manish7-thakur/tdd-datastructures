package trie

type TrieNode struct {
	chars     map[rune]*TrieNode
	endOfWord bool
}

type Trie struct {
	first TrieNode
}

func (t *Trie) Insert(s string) {
	var curr = &t.first
	for _, r := range s {
		if curr.chars == nil {
			curr.chars = make(map[rune]*TrieNode)
		}
		node, ok := curr.chars[r]
		if !ok {
			//Insert into current node
			node = &TrieNode{}
			curr.chars[r] = node
		}
		curr = curr.chars[r]
	}
	curr.endOfWord = true
}

func (t *Trie) Contains(str string, prefixSearch bool) bool {
	next := &t.first
	for _, r := range str {
		node, ok := next.chars[r]
		if !ok {
			return false
		}
		next = node
	}
	return prefixSearch || next.endOfWord
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
		stacklen := len(stack)
		node := stack[stacklen-1]
		r := charStack[stacklen-1]

		child := node.chars[r]
		if len(child.chars) == 0 && !child.endOfWord {
			delete(node.chars, r)
		}
		if len(node.chars) == 0 {
			node = nil
		}
		stack = stack[:stacklen-1]
	}
}
