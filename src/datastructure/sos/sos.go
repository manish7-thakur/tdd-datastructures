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
	if len(n) == 0 {
		return false
	}
	mem := make([][]bool, len(n))
	mem[0] = make([]bool, m+1)
	//first column with sum 0 is true as zero can be made from any set
	for i := 0; i < len(n); i++ {
		mem[i][0] = true
	}
	for i := 0; i < len(n); i++ {
		for j := 1; j <= m; j++ {
			mem[i][j] = mem[i][j-n[i]]
		}
	}
	return mem[len(n)-1][m]
}
