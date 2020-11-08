package tsp

import "math"

func findMin(costs []float64) float64 {
	min := math.MaxFloat64
	for _, v := range costs {
		if v < min {
			min = v
		}
	}
	return min
}

func tsp(costs [][]float64, src int, curr int, vertices []int) float64 {
	if len(vertices) == 1 {
		return costs[curr][vertices[0]] + costs[vertices[0]][src]
	} else {
		calCosts := make([]float64, len(vertices))
		for i, v := range vertices {
			temp := make([]int, len(vertices)-1)
			copy(temp[0:i], vertices[0:i])
			copy(temp[i:], vertices[i+1:])
			calCosts[i] = costs[curr][v] + tsp(costs, src, v, temp)
		}
		return findMin(calCosts)
	}
}

