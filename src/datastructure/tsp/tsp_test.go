package tsp

import (
	"testing"
)

func TestSingleVertexShortestPath(t *testing.T) {
	paths := make([][]float64, 1)
	paths[0] = make([]float64, 1)
	res := tsp(paths, 0, 0, []int{0})
	if res != 0 {
		t.Errorf("expected 0 but found %f", res)
	}
}

func TestTwoVertexShortestPath(t *testing.T) {
	paths := make([][]float64, 2)
	for i := range paths {
		paths[i] = make([]float64, 2)
	}
	paths[0][0] = 0
	paths[0][1] = 2
	paths[1][0] = 3
	paths[1][1] = 0
	res := tsp(paths, 0, 0, []int{0, 1})
	expected := 5.0
	if res != expected {
		t.Errorf("expected %f but found %f", expected, res)
	}
}

func TestThreeVertexShortestPath(t *testing.T) {
	paths := [][]float64{
		{0, 2, 4},
		{3, 0, 4},
		{5, 3, 0}}
	res := tsp(paths, 0, 0, []int{1, 2})
	expected := 10.0
	if res != expected {
		t.Errorf("expected %f but found %f", expected, res)
	}
}

func TestFourVertexShortestPath(t *testing.T) {
	paths := [][]float64{
		{0, 2, 4, 1},
		{3, 0, 2, 3},
		{1, 3, 0, 2},
		{5, 3, 1, 0}}
	res := tsp(paths, 0, 0, []int{1, 2, 3})
	expected := 7.0
	if res != expected {
		t.Errorf("expected %f but found %f", expected, res)
	}
}
