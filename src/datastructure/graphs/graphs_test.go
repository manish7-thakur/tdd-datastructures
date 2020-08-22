package graphs

import (
	"math"
	"reflect"
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
	graph := WeightedGraph{map[string][]Node{}, []string{}}
	res := graph.dijkstraShortestPath(Node{})
	if len(res) != 0 {
		t.Errorf("Expected empty but found %d", len(res))
	}
}

func TestDijkstraTwoNodes(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}, []string{"u", "v"}}
	source := Node{vertex: "u"}
	//Initial distance is INF(MaxInt32) for all nodes except source which has 0
	graph.insert(source.vertex, []Node{{"v", 8, math.MaxInt32, ""}})
	actual := graph.dijkstraShortestPath(source)
	expected := []Path{{"u", 0, ""}, {"v", 8, "u"}}
	if !equal(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestDijkstraThreeNodes(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}, []string{"a", "b", "c"}}
	source := Node{vertex: "a"}
	graph.insert(source.vertex, []Node{{"b", 8, math.MaxInt32, ""}})
	graph.insert("b", []Node{{"c", 5, math.MaxInt32, ""}})
	actual := graph.dijkstraShortestPath(source)
	expected := []Path{{"a", 0, ""}, {"b", 8, "a"}, {"c", 13, "b"}}
	if !equal(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestDijkstraFourNodesNodes(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}, []string{"a", "b", "c", "d"}}
	source := Node{vertex: "a"}
	graph.insert(source.vertex, []Node{{"b", 2, math.MaxInt32, ""}, {"c", 4, math.MaxInt32, ""}, {"d", 1, math.MaxInt32, ""}})
	graph.insert("b", []Node{{"a", 2, math.MaxInt32, ""}, {"c", 1, math.MaxInt32, ""}})
	graph.insert("c", []Node{{"a", 4, math.MaxInt32, ""}, {"b", 1, math.MaxInt32, ""}})
	graph.insert("d", []Node{{"a", 1, math.MaxInt32, ""}})
	actual := graph.dijkstraShortestPath(source)
	expected := []Path{{"a", 0, ""}, {"d", 1, "a"}, {"b", 2, "a"}, {"c", 3, "b"}}
	if !equal(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphDijkstraMultipleVertices(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}, []string{"a", "b", "c", "d", "e", "f", "g"}}
	source := Node{vertex: "a"}
	graph.insert("a", []Node{{"b", 3, math.MaxInt32, ""}, {"c", 2, math.MaxInt32, ""}})
	graph.insert("b", []Node{{"a", 6, math.MaxInt32, ""}, {"c", 7, math.MaxInt32, ""}, {"d", 1, math.MaxInt32, ""}})
	graph.insert("c", []Node{{"a", 2, math.MaxInt32, ""}, {"b", 4, math.MaxInt32, ""}, {"d", 5, math.MaxInt32, ""}})
	graph.insert("d", []Node{{"b", 9, math.MaxInt32, ""}, {"c", 4, math.MaxInt32, ""}, {"e", 8, math.MaxInt32, ""}, {"f", 4, math.MaxInt32, ""}})
	graph.insert("e", []Node{{"d", 4, math.MaxInt32, ""}, {"g", 3, math.MaxInt32, ""}})
	graph.insert("f", []Node{{"d", 6, math.MaxInt32, ""}})
	graph.insert("g", []Node{{"e", 3, math.MaxInt32, ""}})
	var actual = graph.dijkstraShortestPath(source)
	expected := []Path{{"f", 8, "d"}, {"g", 15, "e"}, {"a", 0, ""}, {"b", 3, "a"}, {"c", 2, "a"}, {"d", 4, "b"}, {"e", 12, "d"}}
	if !equal(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphBellmenEmptyGraph(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}, []string{}}
	source := "a"
	var elems, _ = graph.bellmenFord(source)
	var actual [1]Path
	copy(actual[:], elems)
	expected := [1]Path{}
	if actual != expected {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphOneEdgeTwoVertices(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}, []string{"a", "b"}}
	source := "a"
	graph.insert("a", []Node{{"b", -1, 1, "a"}})
	var actual, _ = graph.bellmenFord(source)
	expected := []Path{{"a", 0, ""}, {"b", -1, "a"}}
	if !equal(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphEightEdgesFiveVerticesNegativeWeights(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}, []string{"a", "b", "c", "d", "e"}}
	source := "a"
	graph.insert("a", []Node{{"b", -1, math.MaxInt32, "a"}, {"c", 4, math.MaxInt32, "a"}})
	graph.insert("b", []Node{{"c", 3, math.MaxInt32, "b"}, {"d", 2, math.MaxInt32, "b"}, {"e", 2, math.MaxInt32, "b"}})
	graph.insert("d", []Node{{"c", 5, math.MaxInt32, "d"}, {"b", 1, math.MaxInt32, "d"}})
	graph.insert("e", []Node{{"d", -3, math.MaxInt32, "e"}})
	var actual, _ = graph.bellmenFord(source)
	expected := []Path{{"a", 0, ""}, {"b", -1, "a"}, {"c", 2, "b"}, {"d", -2, "e"}, {"e", 1, "b"}}
	if !equal(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphNegativeCycles(t *testing.T) {
	graph := WeightedGraph{map[string][]Node{}, []string{"a", "b", "c", "d", "e"}}
	source := "a"
	graph.insert("a", []Node{{"b", 1, math.MaxInt32, "a"}})
	graph.insert("b", []Node{{"c", 1, math.MaxInt32, "b"}})
	graph.insert("c", []Node{{"d", -2, math.MaxInt32, "c"}, {"e", 3, math.MaxInt32, "c"}})
	graph.insert("d", []Node{{"b", -1, math.MaxInt32, "d"}})
	var _, err = graph.bellmenFord(source)
	if err.Error() != "-ve cycle found" {
		t.Errorf("Expected to result in error")
	}
}

func equal(actual []Path, expected []Path) bool {
	mapActual := make(map[string]Path)
	mapExpected := make(map[string]Path)
	for _, path := range actual {
		mapActual[path.vertex] = path
	}
	for _, path := range expected {
		mapExpected[path.vertex] = path
	}
	return reflect.DeepEqual(mapActual, mapExpected)
}
