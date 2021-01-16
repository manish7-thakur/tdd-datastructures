package sparsematmul

func multiply(a [3][3]int, b [3][3]int) [3][3]int {
	res := [3][3]int{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			for k := 0; k < len(a); k++ {
				if a[i][k] != 0 && b[k][j] != 0 {
					res[i][j] += a[i][k] * b[k][j]
				}
			}
		}
	}
	return res
}

func multiply2(A [3][3]int, B [3][3]int) [3][3]int {
	res := [len(A)][len(B[0])]int{}
	rowA := [len(A)]bool{}
	colB := [len(B[0])]bool{}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if A[i][j] != 0 {
				rowA[i] = true
				break
			}
		}
	}
	for j := 0; j < len(B[0]); j++ {
		for i := 0; i < len(B); i++ {
			if B[i][j] != 0 {
				colB[j] = true
				break
			}
		}
	}
	for i := 0; i < len(A); i++ {
		for k := 0; k < len(B[0]); k++ {
			if !rowA[i] || !colB[k] {
				res[i][k] = 0
				continue
			}

			sum := 0
			for j := 0; j < len(A[0]); j++ {
				sum += A[i][j] * B[j][k]
			}
			res[i][k] = sum
		}
	}
	return res
}
