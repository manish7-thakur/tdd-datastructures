package rabinkarp

func hash(str string) int {
	var hash int32 = 0
	for _, v := range str {
		hash = 31*hash + v
	}
	return int(hash)
}
