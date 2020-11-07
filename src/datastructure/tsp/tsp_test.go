package tsp

import (
	"math"
	"testing"
)

func TestSingleVertexShortestPath(t *testing.T) {
	paths := make(map[string]float64, 1)
	paths["AA"] = 0
	res := tsp(paths, "A", "A", []string{"A"})
	if res != 0 {
		t.Errorf("expected 0 but found %f", res)
	}
}

func TestTwoVertexShortestPath(t *testing.T) {
	paths := make(map[string]float64, 3)
	paths["AA"] = 0
	paths["AB"] = 2
	paths["BA"] = 3
	res := tsp(paths, "A", "A", []string{"B"})
	expected := 5.0
	if res != expected {
		t.Errorf("expected %f but found %f", expected, res)
	}
}

func TestThreVertexShortestPath(t *testing.T) {
	paths := make(map[string]float64, 6)
	paths["AA"] = 0
	paths["AB"] = 2
	paths["AC"] = 4
	paths["BA"] = 3
	paths["BB"] = 0
	paths["BC"] = 4
	paths["CA"] = 5
	paths["CB"] = 3
	paths["CC"] = 0
	res := tsp(paths, "A", "A", []string{"B", "C"})
	expected := 10.0
	if res != expected {
		t.Errorf("expected %f but found %f", expected, res)
	}
}

func tsp(costs map[string]float64, src string, curr string, vertices []string) float64 {
	if len(vertices) == 1 {
		return costs[curr+vertices[0]] + costs[vertices[0]+src]
	} else {
		v1 := costs[src+vertices[0]] + tsp(costs, src, vertices[0], vertices[1:])
		v2 := costs[src+vertices[1]] + tsp(costs, src, vertices[1], vertices[:1])
		return math.Min(v1, v2)
	}
}
