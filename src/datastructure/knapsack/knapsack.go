package knapsack

import "datastructure/math"

func MaxProfit(weights []int, profits []int, capacity int) int {
	weightlen := len(weights)
	if weightlen == 0 || capacity == 0 {
		return 0
	}
	if weights[weightlen-1] > capacity {
		return MaxProfit(weights[0:weightlen-1], profits[0:weightlen-1], capacity)
	}
	return math.Max(profits[weightlen-1]+MaxProfit(weights[0:weightlen-1], profits[0:weightlen-1], capacity-weights[weightlen-1]),
		MaxProfit(weights[0:weightlen-1], profits[0:weightlen-1], capacity))
}

func MaxProfitIter(weights []int, profits []int, capacity int) int {
	weightlen := len(weights)
	if weightlen == 0 || capacity == 0 {
		return 0
	}
	mem := make([][]int, weightlen)
	for i := 0; i < weightlen; i++ {
		mem[i] = make([]int, capacity+1)
	}
	for j := 0; j <= capacity; j++ {
		if j < weights[0] {
			mem[0][j] = 0
		} else {
			mem[0][j] = profits[j]
		}

	}
	return mem[0][capacity]
}
