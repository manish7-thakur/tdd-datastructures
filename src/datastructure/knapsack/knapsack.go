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
	if len(weights) == 0 || capacity == 0 {
		return 0
	}
	return profits[0]
}
