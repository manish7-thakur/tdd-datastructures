package lcs

func LCS(s1 string, s2 string) int {
	return lcsCount(s1, s2, len(s1), len(s2))
}

func lcsCount(s1 string, s2 string, i int, j int) int {
	if i < 1 || j < 1 {
		return 0
	} else if s1[i-1] == s2[j-1] {
		return 1 + lcsCount(s1, s2, i-1, j-1)
	} else {
		return Max(lcsCount(s1, s2, i-1, j), lcsCount(s1, s2, i, j-1))
	}
}
func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
