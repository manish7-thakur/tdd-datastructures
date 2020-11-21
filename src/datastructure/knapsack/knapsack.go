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
	for i := 0; i < weightlen; i++ {
		for j := 0; j <= capacity; j++ {
			if j < weights[i] {
				if i > 0 {
					mem[i][j] = mem[i-1][j]
				}
			} else {
				if i > 0 {
					mem[i][j] = math.Max(profits[i]+mem[i-1][j-weights[i]], mem[i-1][j])
				} else {
					mem[i][j] = profits[i]
				}
			}
		}
	}
	return mem[weightlen-1][capacity]
}
