package graphs

import (
	"math"
	"testing"
)

func TestGraphInsertVertex(t *testing.T) {
	graph := Graph{map[string][]string{}}
	graph.insert("a", "b")
	if len(graph.adjList) != 1 {
		t.Errorf("Expected 1 vertex but found %d", len(graph.adjList))
	}
}

func TestGraphAppendToExistingList(t *testing.T) {
	graph := Graph{map[string][]string{}}
	vertex := "a"
	graph.insert(vertex, "b")
	graph.insert(vertex, "c")
	var expected [2]string
	copy(expected[:], graph.adjList[vertex])
	if expected != [2]string{"b", "c"} {
		t.Errorf("Expected 1 vertex but found %d", len(graph.adjList))
	}
}

func TestGraphBFSEmpty(t *testing.T) {
	graph := Graph{map[string][]string{}}
	elems := graph.bfs("")
	var actual [1]string
	copy(actual[:], elems)
	expected := [1]string{}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphBFSTwoVertices(t *testing.T) {
	graph := Graph{map[string][]string{}}
	graph.insert("u", "v")
	var elems = graph.bfs("u")
	var actual [2]string
	copy(actual[:], elems)
	expected := [2]string{"u", "v"}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphBFSMultipleVertices(t *testing.T) {
	graph := Graph{map[string][]string{}}
	graph.insertAdj("a", []string{"b", "c"})
	graph.insertAdj("b", []string{"a", "c", "d"})
	graph.insertAdj("c", []string{"a", "b", "d"})
	graph.insertAdj("d", []string{"b", "c", "e", "f"})
	graph.insertAdj("e", []string{"d", "g"})
	graph.insertAdj("f", []string{"d"})
	graph.insertAdj("g", []string{"e"})
	var elems = graph.bfs("a")
	var actual [7]string
	copy(actual[:], elems)
	expected := [7]string{"a", "b", "c", "d", "e", "f", "g"}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphBFSMultipleVerticesLastNode(t *testing.T) {
	graph := Graph{map[string][]string{}}
	graph.insertAdj("a", []string{"b", "c"})
	graph.insertAdj("b", []string{"a", "c", "d"})
	graph.insertAdj("c", []string{"a", "b", "d"})
	graph.insertAdj("d", []string{"b", "c", "e", "f"})
	graph.insertAdj("e", []string{"d", "g"})
	graph.insertAdj("f", []string{"d"})
	graph.insertAdj("g", []string{"e"})
	var elems = graph.bfs("g")
	var actual [7]string
	copy(actual[:], elems)
	expected := [7]string{"g", "e", "d", "b", "c", "f", "a"}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphDFSEmpty(t *testing.T) {
	graph := Graph{map[string][]string{}}
	vertices := graph.dfs("")
	expected := [1]string{}
	var actual [1]string
	copy(actual[:], vertices)
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphDFSTwoVertices(t *testing.T) {
	graph := Graph{map[string][]string{}}
	graph.insertAdj("u", []string{"v"})
	vertices := graph.dfs("u")
	var actual [2]string
	copy(actual[:], vertices)
	expected := [2]string{"u", "v"}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphDFSMultipleVertices(t *testing.T) {
	graph := Graph{map[string][]string{}}
	graph.insertAdj("a", []string{"b", "c"})
	graph.insertAdj("b", []string{"a", "c", "d"})
	graph.insertAdj("c", []string{"a", "b", "d"})
	graph.insertAdj("d", []string{"b", "c", "e", "f"})
	graph.insertAdj("e", []string{"d", "g"})
	graph.insertAdj("f", []string{"d"})
	graph.insertAdj("g", []string{"e"})
	var elems = graph.dfs("a")
	var actual [7]string
	copy(actual[:], elems)
	expected := [7]string{"a", "c", "d", "f", "e", "g", "b"}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphDFSMultipleVerticesLastNode(t *testing.T) {
	graph := Graph{map[string][]string{}}
	graph.insertAdj("a", []string{"b", "c"})
	graph.insertAdj("b", []string{"a", "c", "d"})
	graph.insertAdj("c", []string{"a", "b", "d"})
	graph.insertAdj("d", []string{"b", "c", "e", "f"})
	graph.insertAdj("e", []string{"d", "g"})
	graph.insertAdj("f", []string{"d"})
	graph.insertAdj("g", []string{"e"})
	var elems = graph.dfs("g")
	var actual [7]string
	copy(actual[:], elems)
	expected := [7]string{"g", "e", "d", "f", "c", "b", "a"}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestDijkstraEmptyGraph(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}}
	res := graph.dijkstraShortestPath(Node{})
	if len(res) != 1 {
		t.Errorf("Expected empty but found %d", len(res))
	}
}

func TestDijkstraTwoNodes(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}}
	source := Node{vertex: "u"}
	graph.insert(source.vertex, []Node{{"v", 8, math.MaxInt32}})
	res := graph.dijkstraShortestPath(source)
	var actual [2]Node
	copy(actual[:], res)
	expected := [2]Node{{"u", 0, 0}, {"v", 8, 8}}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestDijkstraThreeNodes(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}}
	source := Node{vertex: "a"}
	graph.insert(source.vertex, []Node{{"b", 8, math.MaxInt32}})
	graph.insert("b", []Node{{"c", 5, math.MaxInt32}})
	res := graph.dijkstraShortestPath(source)
	var actual [3]Node
	copy(actual[:], res)
	expected := [3]Node{{"a", 0, 0}, {"b", 8, 8}, {"c", 5, 13}}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}
