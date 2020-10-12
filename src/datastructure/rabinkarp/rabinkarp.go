package rabinkarp

import "math"

func New(str string) HashGen {
	priorHash := Hash(str)
	return HashGen{priorHash, len(str), math.Pow(31, float64(len(str)))}
}

type HashGen struct {
	hash           int
	strlen         int
	multiplicative float64
}

func (g *HashGen) RollHash(next, old byte) {
	g.hash = newHash(g.hash, next, old, int(g.multiplicative))
}

func Hash(str string) int {
	var hash int32 = 0
	for _, v := range str {
		hash = 31*hash + v
	}
	return int(hash)
}

func newHash(priorHash int, nextChar byte, oldChar byte, multiplicative int) int {
	hash := (31*priorHash + int(nextChar)) - (multiplicative * int(oldChar))
	return hash
}

func IndexRabinKarp(s, substr string) int {
	substrlen := len(substr)
	slen := len(s)
	if substrlen == 0 {
		return 0
	} else if substrlen > slen {
		return -1
	} else {
		//compute hash substring
		//match with hash from string s part
		//if matches compare string else slide
		//if char matches return current idx
		hash := Hash(substr)
		hashGen := New(s[:substrlen])
		for i := 0; i <= slen-substrlen; i++ {
			if hashGen.hash == hash {
				if substr == s[i:i+substrlen] {
					return i
				}
			}
			if i+substrlen < slen {
				hashGen.RollHash(s[i+substrlen], s[i])
			}
		}
	}
	return -1
}
