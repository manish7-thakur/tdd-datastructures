package rabinkarp

func hash(str string) int {
	var hash int32 = 0
	for _, v := range str {
		hash = 31*hash + v
	}
	return int(hash)
}

func NewHash(priorHash int, nextChar string, oldChar string) int {
	hash := ((31*priorHash + int(nextChar[0])) - int(oldChar[0])) / 31
	return hash
}
