package kmp

func KMPSearch(text string, pattern string) int {
	patlen := len(pattern)
	textlen := len(text)
	if patlen == 0 {
		return 0
	} else if patlen > textlen {
		return -1
	}
	j := 0
	lps := LPS(pattern)
	for i := 0; i < textlen; i++ {
		if pattern[j] == text[i] {
			j++
		} else {
			for ; j > 0 && pattern[j] != text[i]; {
				j = lps[j-1]
			}
			if pattern[j] == text[i] {
				j++
			}
		}
		if j == patlen {
			return i - patlen + 1
		}
	}
	return -1
}

func LPS(str string) []int {
	strlen := len(str)
	if strlen == 0 {
		return []int{}
	}
	lpsarr := make([]int, strlen)
	j := 0
	for i := 1; i < strlen; i++ {
		if str[i] == str[j] {
			lpsarr[i] = j + 1
			j++
		} else {
			for ; j > 0 && str[i] != str[j]; {
				j = lpsarr[j-1]
			}
			if str[i] == str[j] {
				lpsarr[i] = j + 1
				j++
			}
		}
	}
	return lpsarr
}
