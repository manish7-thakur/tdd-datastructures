package max_cost_path

import (
	"math"
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
	var mat [][]float64
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

func computeMaxCost(mat [][]float64, i int, j int) float64 {
	if i < 0 || j < 0 {
		return 0
	} else {
		return mat[i][j] +
			math.Max(math.Max(computeMaxCost(mat, i-1, j-1), computeMaxCost(mat, i, j-1)),
				computeMaxCost(mat, i-1, j))
	}
}
