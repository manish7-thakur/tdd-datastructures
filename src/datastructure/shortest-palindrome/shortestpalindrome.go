package shortest_palindrome

import (
	"datastructure/kmp"
	"strings"
)

func ShortestPalindrome(s string) string {
	strlen := len(s)
	if strlen < 2 {
		return s
	}
	var b strings.Builder
	b.Grow(2 * strlen)
	b.WriteString(s)
	b.WriteByte('#')
	for i := strlen - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	str := b.String()
	kmp_arr := kmp.LPS(str)
	arraylen := len(kmp_arr)
	idxvalue := kmp_arr[arraylen-1]
	newstrlen := len(str)
	prefix := str[strings.IndexByte(str, '#')+1 : newstrlen-idxvalue]
	b.Reset()
	prefixlen := len(prefix)
	b.Grow(prefixlen + strlen)
	println(b.Cap())
	b.WriteString(prefix)
	b.WriteString(s)
	return b.String()
}
