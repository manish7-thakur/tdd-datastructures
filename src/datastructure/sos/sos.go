package sos

func sos(m int, n []int) bool {
	if m == 0 {
		return true
	}
	if len(n) == 0 {
		return false
	}
	return sos(m-n[0], n[1:]) || sos(m, n[1:])
}
