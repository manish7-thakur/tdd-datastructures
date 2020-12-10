package misc

import (
	"strconv"
)

func MaxNum(s string) int {
	strlen := len(s)
	if strlen == 0 {
		return 0
	}
	max := -1
	oneDigitVal := -1
	twoDigitVal := -1
	for i := 0; i < strlen; i++ {
		oneDigitVal, _ = strconv.Atoi(s[i : i+1])
		if i+1 < strlen {
			twoDigitVal, _ = strconv.Atoi(s[i : i+2])
		}
		if oneDigitVal > twoDigitVal {
			if oneDigitVal > max {
				max = oneDigitVal
			}
		} else {
			if twoDigitVal > max {
				max = twoDigitVal
			}
		}
	}
	return max
}
