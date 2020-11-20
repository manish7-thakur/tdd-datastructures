package knapsack

import "testing"

func TestKnapsackEmptyWeights(t *testing.T) {
	capacity := 5
	res := MaxProfit([]int{}, []int{}, capacity)
	if res != 0 {
		t.Errorf("Expected empty but found %d", res)
	}
}

func TestKnapsackEmptyWeightsIter(t *testing.T) {
	capacity := 5
	res := MaxProfitIter([]int{}, []int{}, capacity)
	if res != 0 {
		t.Errorf("Expected empty but found %d", res)
	}
}

func TestKnapsackWeightsZeroCapacity(t *testing.T) {
	capacity := 0
	res := MaxProfit([]int{}, []int{}, capacity)
	if res != 0 {
		t.Errorf("Expected empty but found %d", res)
	}
}

func TestKnapsackWeightsZeroCapacityIter(t *testing.T) {
	capacity := 0
	res := MaxProfitIter([]int{}, []int{}, capacity)
	if res != 0 {
		t.Errorf("Expected empty but found %d", res)
	}
}

func TestKnapsackSingleWeight(t *testing.T) {
	capacity := 5
	weights := []int{5}
	profits := []int{3}
	res := MaxProfit(weights, profits, capacity)
	expected := 3
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestKnapsackSingleWeightIter(t *testing.T) {
	capacity := 5
	weights := []int{5}
	profits := []int{3}
	res := MaxProfitIter(weights, profits, capacity)
	expected := 3
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestKnapsackSingleWeightHigherThanCapacity(t *testing.T) {
	capacity := 4
	weights := []int{5}
	profits := []int{3}
	res := MaxProfit(weights, profits, capacity)
	expected := 0
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestKnapsackSingleWeightHigherThanCapacityIter(t *testing.T) {
	capacity := 1
	weights := []int{5}
	profits := []int{3}
	res := MaxProfitIter(weights, profits, capacity)
	expected := 0
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestKnapsackDoubleWeights(t *testing.T) {
	capacity := 6
	weights := []int{6, 5}
	profits := []int{4, 3}
	res := MaxProfit(weights, profits, capacity)
	expected := 4
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestKnapsackDoubleWeightsIter(t *testing.T) {
	capacity := 2
	weights := []int{1, 2}
	profits := []int{3, 4}
	res := MaxProfitIter(weights, profits, capacity)
	expected := 4
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestKnapsackDoubleWeightsBothIncludedIter(t *testing.T) {
	capacity := 2
	weights := []int{1, 1}
	profits := []int{3, 4}
	res := MaxProfitIter(weights, profits, capacity)
	expected := 7
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestKnapsackTripleWeights(t *testing.T) {
	capacity := 5
	weights := []int{3, 4, 2}
	profits := []int{1, 3, 3}
	res := MaxProfit(weights, profits, capacity)
	expected := 4
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestKnapsackTripleWeightsIter(t *testing.T) {
	capacity := 3
	weights := []int{1, 4, 2}
	profits := []int{1, 3, 3}
	res := MaxProfitIter(weights, profits, capacity)
	expected := 4
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestKnapsackFourWeights(t *testing.T) {
	capacity := 7
	weights := []int{1, 3, 4, 5}
	profits := []int{1, 4, 5, 7}
	res := MaxProfit(weights, profits, capacity)
	expected := 9
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}

func TestKnapsackMultipleWeights(t *testing.T) {
	capacity := 84
	weights := []int{15, 34, 43, 45, 35, 23, 4, 7, 3, 18, 2, 29, 2, 5, 3, 67}
	profits := []int{14, 54, 56, 78, 2, 53, 62, 12, 34, 2, 45, 23, 5, 2, 4, 3}
	res := MaxProfit(weights, profits, capacity)
	expected := 284
	if res != expected {
		t.Errorf("Expected %d but found %d", expected, res)
	}
}
