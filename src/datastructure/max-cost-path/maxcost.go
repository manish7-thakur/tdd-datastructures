package max_cost_path

import "math"

func computeMaxCost(mat [][]float64, i int, j int) float64 {
	if i < 0 || j < 0 {
		return 0
	} else {
		return mat[i][j] +
			math.Max(math.Max(computeMaxCost(mat, i-1, j-1), computeMaxCost(mat, i, j-1)),
				computeMaxCost(mat, i-1, j))
	}
}
