package sparsematmul

func multiply(a [3][3]int, b [3][3]int) [3][3]int {
	res := [3][3]int{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			for k := 0; k < len(a); k++ {
				if(a[i][k] !=0 && b[k][j] != 0) {
					res[i][j] += a[i][k] * b[k][j];
				}
			}
		}
	}
	return res
}