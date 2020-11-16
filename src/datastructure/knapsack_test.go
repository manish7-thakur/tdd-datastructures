package datastructure

import "testing"

func TestKnapsackEmptyWeights(t *testing.T) {
	capacity := 5
	res := maxProfit([]int{}, []int{}, capacity)
	if res != 0 {
		t.Errorf("Expected empty but found %d", res)
	}
}

func TestKnapsackWeightsZeroCapacity(t *testing.T) {
	capacity := 0
	res := maxProfit([]int{}, []int{}, capacity)
	if res != 0 {
		t.Errorf("Expected empty but found %d", res)
	}
}

func TestKnapsackSingleWeight(t *testing.T) {
	capacity := 5
	weights := []int{5}
	profits := []int{3}
	res := maxProfit(weights, profits, capacity)
	expected := 3
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestKnapsackSingleWeightHigherThanCapacity(t *testing.T) {
	capacity := 4
	weights := []int{5}
	profits := []int{3}
	res := maxProfit(weights, profits, capacity)
	expected := 0
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func maxProfit(weights []int, profits []int, capacity int) int {
	weightlen := len(weights)
	if weightlen == 0 || capacity == 0 {
		return 0
	}
	if weights[weightlen-1] > capacity {
		return maxProfit(weights[0:weightlen-1], profits[0:weightlen-1], capacity)
	} else {
		return profits[weightlen-1] + maxProfit(weights[0:weightlen-1], profits[0:weightlen-1], capacity-weights[weightlen-1])
	}
}
