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
	setlen := len(n)
	if setlen == 0 {
		return false
	}
	mem := make([][]bool, setlen)
	//first column with sum 0 is true as zero can be made from any set
	for i := 0; i < setlen; i++ {
		mem[i] = make([]bool, m+1)
		mem[i][0] = true
	}
	for i := 0; i < setlen; i++ {
		for j := 1; j <= m; j++ {
			if j == n[i] {
				mem[i][j] = true
			} else if j > n[i] && i > 0 {
				//use the above value if its true
				mem[i][j] = mem[i-1][j] || mem[i-1][j-n[i]]
			} else if j < n[i] && i > 0 {
				//copy from top until j == n[i]
				mem[i][j] = mem[i-1][j]
			}
		}
	}
	return mem[setlen-1][m]
}
