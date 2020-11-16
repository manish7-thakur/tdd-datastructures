package datastructure

import "testing"

func TestKnapsackEmptyWeights(t *testing.T) {
	capacity := 5
	res := maxProfit([]int{}, []int{}, capacity, 0)
	if res != 0 {
		t.Errorf("Expected empty but found %d", res)
	}
}

func TestKnapsackSingleWeight(t *testing.T) {
	capacity := 5
	weights := []int{5}
	profits := []int{3}
	res := maxProfit(weights, profits, capacity, 0)
	expected := 3
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestKnapsackSingleWeightHigherThanCapacity(t *testing.T) {
	capacity := 4
	weights := []int{5}
	profits := []int{3}
	res := maxProfit(weights, profits, capacity, 0)
	expected := 0
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func maxProfit(weights []int, profits []int, capacity int, currProfit int) int {
	weightlen := len(weights)
	if weightlen != 0 {
		if weights[weightlen-1] > capacity {
			maxProfit(weights[0:weightlen-1], profits[0:weightlen-1], capacity, currProfit)
		} else {
			currProfit += maxProfit(weights[0:weightlen-1], profits[0:weightlen-1], capacity-weights[weightlen-1], profits[weightlen-1])
		}
	}
	return currProfit
}
