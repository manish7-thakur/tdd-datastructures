package misc

import (
	"fmt"
)

/*
Given a string S consisting of N lowercase letters, return the minimum number of letters
that must be deleted to obtain a word in which every letter occurs a unique number of times.We only care
about occurrences of letters that appear at least once in result.

E.g
1. Given S = "aaaabbbb", the function should return 1. We can delete one occurrence of a or one occurrence of b.
Then one letter will occur 4 times and one 3 times.

2. Given S = "ccaaffddeccee", the function should return 6. For e.g. we can delete all occurrences of e & f and one occurrence of d to
obtain the word "ccaadc". Note that both e & f will occur zero times in the new word, but that is fine, since we only care about letters
that appear at-least once.

3. Given S = "eeee", the function should return 0 (there is no need to delete any characters).

4. Given S = "example", the function should return 4.

Write an efficient algorithm for the following assumptions:
	*N is an integer within the range [0....300,000]
	*string S consists only of lowercase letters (a-z)

*/

func MinCharsToDelete(s string) int {
	occMap := FindOccurrence(s)
	return DeleteMinCharsMap(occMap)
}

func FindOccurrence(s string) map[int]int {
	charMap := make(map[rune]int)
	occMap := make(map[int]int)
	for _, c := range s {
		if oldVal, ok := occMap[charMap[c]]; ok {
			if oldVal == 1 {
				delete(occMap, charMap[c])
			} else {
				occMap[charMap[c]]--
			}
		}
		charMap[c]++
		occMap[charMap[c]]++
	}
	return occMap
}

func DeleteMinCharsMap(arr map[int]int) int {
	total := 0
	for k, v := range arr {
		if _, ok := arr[k-1]; ok && v > 1 {
			total += (v - 1) * k
		} else if v > 1 {
			total += 1 + (v-2)*k
		}
	}
	return total
}

func pattern(n int) {
	len := 1
	for i := 2; i <= n; i *= 2 {
		for j := 0; j < i; j++ {
			fmt.Printf("%0*b\n", len, j)
		}
		len++
	}
}
