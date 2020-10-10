package rabinkarp

import "math"

func New(str string) HashGen {
	priorHash := Hash(str)
	return HashGen{priorHash, len(str)}
}

type HashGen struct {
	hash   int
	strlen int
}

func (g HashGen) RollHash(next int32, old int32) int {
	g.hash = newHash(g.hash, next, old, float64(g.strlen))
	return g.hash
}

func Hash(str string) int {
	var hash int32 = 0
	for _, v := range str {
		hash = 31*hash + v
	}
	return int(hash)
}

func newHash(priorHash int, nextChar rune, oldChar rune, strlen float64) int {
	hash := (31*priorHash + int(nextChar)) - (int(math.Pow(31, strlen)) * int(oldChar))
	return hash
}
