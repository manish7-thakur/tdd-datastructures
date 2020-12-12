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
	mid := strlen / 2
	i := 0
	j := 0
	if strlen%2 == 0 {
		i = mid - 1
		j = mid
	} else {
		i = mid
		j = mid
	}
	prefixarr := kmp.LPS(s)
	for ; i >= 0; {
		if s[i] == s[j] {
			i--
			j++
		} else {
			for ; i != j && s[i] != s[j]; {
				if prefixarr[j-1] > 0 {
					j = prefixarr[j-1]
				} else {
					j--
				}
			}
		}
	}
	var b strings.Builder
	b.Grow(2*strlen - j)
	for i = strlen - 1; i >= j; i-- {
		b.WriteByte(s[i])
	}
	b.WriteString(s)
	return b.String()
}
