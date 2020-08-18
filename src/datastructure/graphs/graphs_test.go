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
	//Initial distance is INF(MaxInt32) for all nodes except source which has 0
	graph.insert(source.vertex, []Node{{"v", 8, math.MaxInt32, ""}})
	res := graph.dijkstraShortestPath(source)
	var actual [2]Node
	copy(actual[:], res)
	expected := [2]Node{{"u", 0, 0, ""}, {"v", 8, 8, ""}}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestDijkstraThreeNodes(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}}
	source := Node{vertex: "a"}
	graph.insert(source.vertex, []Node{{"b", 8, math.MaxInt32, ""}})
	graph.insert("b", []Node{{"c", 5, math.MaxInt32, ""}})
	res := graph.dijkstraShortestPath(source)
	var actual [3]Node
	copy(actual[:], res)
	expected := [3]Node{{"a", 0, 0, ""}, {"b", 8, 8, ""}, {"c", 5, 13, ""}}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestDijkstraFourNodesNodes(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}}
	source := Node{vertex: "a"}
	graph.insert(source.vertex, []Node{{"b", 2, math.MaxInt32, ""}, {"c", 4, math.MaxInt32, ""}, {"d", 1, math.MaxInt32, ""}})
	graph.insert("b", []Node{{"a", 2, math.MaxInt32, ""}, {"c", 1, math.MaxInt32, ""}})
	graph.insert("c", []Node{{"a", 4, math.MaxInt32, ""}, {"b", 1, math.MaxInt32, ""}})
	graph.insert("d", []Node{{"a", 1, math.MaxInt32, ""}})
	res := graph.dijkstraShortestPath(source)
	var actual [4]Node
	copy(actual[:], res)
	expected := [4]Node{{"a", 0, 0, ""}, {"d", 1, 1, ""}, {"b", 2, 2, ""}, {"c", 1, 3, ""}}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphDijkstraMultipleVertices(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}}
	source := Node{vertex: "a"}
	graph.insert("a", []Node{{"b", 3, math.MaxInt32, ""}, {"c", 2, math.MaxInt32, ""}})
	graph.insert("b", []Node{{"a", 6, math.MaxInt32, ""}, {"c", 7, math.MaxInt32, ""}, {"d", 1, math.MaxInt32, ""}})
	graph.insert("c", []Node{{"a", 2, math.MaxInt32, ""}, {"b", 4, math.MaxInt32, ""}, {"d", 5, math.MaxInt32, ""}})
	graph.insert("d", []Node{{"b", 9, math.MaxInt32, ""}, {"c", 4, math.MaxInt32, ""}, {"e", 8, math.MaxInt32, ""}, {"f", 4, math.MaxInt32, ""}})
	graph.insert("e", []Node{{"d", 4, math.MaxInt32, ""}, {"g", 3, math.MaxInt32, ""}})
	graph.insert("f", []Node{{"d", 6, math.MaxInt32, ""}})
	graph.insert("g", []Node{{"e", 3, math.MaxInt32, ""}})
	var elems = graph.dijkstraShortestPath(source)
	var actual [7]Node
	copy(actual[:], elems)
	expected := [7]Node{{"a", 0, 0, ""}, {"c", 2, 2, ""}, {"b", 3, 3, ""}, {"d", 1, 4, ""}, {"f", 4, 8, ""}, {"e", 8, 12, ""}, {"g", 3, 15, ""}}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphBellmenEmptyGraph(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}}
	source := Node{vertex: "a"}
	var elems = graph.bellmenFord(source, 0)
	var actual [1]Node
	copy(actual[:], elems)
	expected := [1]Node{{"", 0, 0, ""}}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphTwoVertices(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}}
	source := Node{vertex: "a"}
	graph.insert("a", []Node{{"b", -1, math.MaxInt32, ""}})
	var elems = graph.bellmenFord(source, 2)
	var actual [1]Node
	copy(actual[:], elems)
	expected := [1]Node{{"b", -1, -1, "a"}}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphThreeVertices(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}}
	source := Node{vertex: "c"}
	graph.insert("a", []Node{{"b", -1, math.MaxInt32, ""}, {"c", 2, math.MaxInt32, ""}, {"d", 1, math.MaxInt32, ""}})
	graph.insert("b", []Node{{"d", 3, math.MaxInt32, ""}})
	graph.insert("c", []Node{{"d", - 1, math.MaxInt32, ""}, {"a", 1, math.MaxInt32, ""}})
	var elems = graph.bellmenFord(source, 4)
	var actual [6]Node
	copy(actual[:], elems)
	expected := [6]Node{{"b", -1, 0, "a"}, {"c", 2, 0, "a"}, {"d", 1, 2, "a"}, {"d", 3, 3, "b"}, {"d", -1, -1, "c"}, {"a", 1, 1, "c"}}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}
