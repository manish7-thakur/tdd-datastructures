package rabinkarp

import "math"

func hash(str string) int {
	var hash int32 = 0
	for _, v := range str {
		hash = 31*hash + v
	}
	return int(hash)
}

func NewHash(priorHash int, nextChar rune, oldChar rune, strlen float64) int {
	hash := (31*priorHash + int(nextChar)) - (int(math.Pow(31, strlen)) * int(oldChar))
	return hash
}
