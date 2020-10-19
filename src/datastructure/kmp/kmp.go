package kmp

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
