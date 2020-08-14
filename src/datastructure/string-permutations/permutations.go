package string_permutations

func permutations(s string) []string {
	var allPerm []string
	if len(s) == 0 {
		allPerm = append(allPerm, s)
		return allPerm
	} else {
		for i := 0; i < len(s); i = i + 1 {
			a := []rune(s)
			current := a[i]
			firstH := a[0:i]
			secondH := a[i+1:]
			shortPer := permutations(string(firstH) + string(secondH))
			for _, perms := range shortPer {
				allPerm = append(allPerm, string(current)+perms)
			}
		}
	}
	return allPerm
}

