package tsp

import (
	"testing"
)

func TestSingleVertexShortestPath(t *testing.T) {
	paths := make(map[string]float64, 1)
	paths["AA"] = 0
	res := tsp(paths, "A", "A", []string{})
	if res != 0 {
		t.Errorf("expected 0 but found %f", res)
	}
}

func TestTwoVertexShortestPath(t *testing.T) {
	paths := make(map[string]float64, 1)
	paths["AA"] = 0
	paths["AB"] = 2
	paths["BA"] = 3
	res := tsp(paths, "A", "A", []string{"B"})
	if res != 5 {
		t.Errorf("expected 0 but found %f", res)
	}
}

func tsp(costs map[string]float64, src string, curr string, vertices []string) float64 {
	if len(vertices) == 0 {
		return costs[curr+src]
	} else {
		v1 := costs[src+vertices[0]] + tsp(costs, src, vertices[0], vertices[1:])
		return v1
	}
}
