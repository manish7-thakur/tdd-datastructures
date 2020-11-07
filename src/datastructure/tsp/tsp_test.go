package tsp

import (
	"math"
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
	paths := make([][]float64, 3)
	for i := range paths {
		paths[i] = make([]float64, 3)
	}
	paths[0][0] = 0
	paths[0][1] = 2
	paths[0][2] = 4
	paths[1][0] = 3
	paths[1][1] = 0
	paths[1][2] = 4
	paths[2][0] = 5
	paths[2][1] = 3
	paths[2][2] = 0
	res := tsp(paths, 0, 0, []int{1, 2})
	expected := 10.0
	if res != expected {
		t.Errorf("expected %f but found %f", expected, res)
	}
}

func TestFourVertexShortestPath(t *testing.T) {
	paths := make([][]float64, 4)
	for i := range paths {
		paths[i] = make([]float64, 4)
	}
	paths[0][0] = 0
	paths[0][1] = 2
	paths[0][2] = 4
	paths[0][3] = 1
	paths[1][0] = 3
	paths[1][1] = 0
	paths[1][2] = 2
	paths[1][3] = 3
	paths[2][0] = 1
	paths[2][1] = 3
	paths[2][2] = 0
	paths[2][3] = 2
	paths[3][0] = 5
	paths[3][1] = 3
	paths[3][2] = 1
	paths[3][3] = 0
	res := tsp(paths, 0, 0, []int{1, 2, 3})
	expected := 7.0
	if res != expected {
		t.Errorf("expected %f but found %f", expected, res)
	}
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

func findMin(costs []float64) float64 {
	min := math.MaxFloat64
	for _, v := range costs {
		if v < min {
			min = v
		}
	}
	return min
}
