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

func TestThreVertexShortestPath(t *testing.T) {
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

func tsp(costs [][]float64, src int, curr int, vertices []int) float64 {
	if len(vertices) == 1 {
		return costs[curr][vertices[0]] + costs[vertices[0]][src]
	} else {
		v1 := costs[src][vertices[0]] + tsp(costs, src, vertices[0], vertices[1:])
		v2 := costs[src][vertices[1]] + tsp(costs, src, vertices[1], vertices[:1])
		return math.Min(v1, v2)
	}
}
