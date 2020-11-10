package sos

func sos(m int, n []int) bool {
	if m == 0 {
		return true
	}
	if len(n) == 0 {
		return false
	}
	if n[0] > m {
		return sos(m, n[1:])
	}
	return sos(m-n[0], n[1:]) || sos(m, n[1:])
}

func sosIter(m int, n []int) bool {
	if m == 0 {
		return true
	}
	var num int
	for i := 0; i < len(n); i++ {
		num = m
		for j := i; j < len(n); j++ {
			if n[j] <= num {
				num -= n[j]
			}
			if num == 0 {
				return true
			}
		}
	}
	return false
}
