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
	idxvalue := kmp_arr[len(kmp_arr)-1]
	b.Reset()
	prefix := str[strings.IndexByte(str, '#')+1 : len(str)-idxvalue]
	b.WriteString(prefix)
	b.WriteString(s)
	return b.String()
}
