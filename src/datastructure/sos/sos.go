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
	for _, num := range n {
		if num <= m {
			m-=num
		}
		if m == 0 {
			return true
		}
	}
	return false
}
