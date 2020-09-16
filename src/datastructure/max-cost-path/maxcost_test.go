package max_cost_path

import (
	"testing"
)

func TestMaxCostRowBoundNegative(t *testing.T) {
	var mat [][]float64
	actual := computeMaxCost(mat, -1, 0)
	expected := 0.0
	if actual != expected {
		t.Errorf("Expected %f but found %f", actual, expected)
	}
}

func TestMaxCostColumnBoundNegative(t *testing.T) {
	var mat [][]float64
	actual := computeMaxCost(mat, 0, -1)
	expected := 0.0
	if actual != expected {
		t.Errorf("Expected %f but found %f", expected, actual)
	}
}

func TestMaxCostEmptyMatrix(t *testing.T) {
	mat := [][]float64{{0}}
	actual := computeMaxCost(mat, 0, 0)
	expected := 0.0
	if actual != expected {
		t.Errorf("Expected %f but found %f", expected, actual)
	}
}

func TestMaxCost2X1Matrix(t *testing.T) {
	mat := [][]float64{{1}, {2}}
	actual := computeMaxCost(mat, 1, 0)
	expected := 3.0
	if actual != expected {
		t.Errorf("Expected %f but found %f", expected, actual)
	}
}
