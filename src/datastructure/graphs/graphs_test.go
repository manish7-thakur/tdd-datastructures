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
	graph := WeightedGraph{map[string][]Edge{}, []string{}}
	res := graph.dijkstraShortestPath(Edge{})
	if len(res) != 0 {
		t.Errorf("Expected empty but found %d", len(res))
	}
}

func TestDijkstraTwoNodes(t *testing.T) {
	graph := WeightedGraph{map[string][]Edge{}, []string{"u", "v"}}
	source := Edge{vertex: "u"}
	//Initial distance is INF(MaxInt32) for all nodes except source which has 0
	graph.insert(source.vertex, []Edge{{"v", 8, ""}})
	actual := graph.dijkstraShortestPath(source)
	expected := []Path{{"u", 0, ""}, {"v", 8, "u"}}
	if !equal(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestDijkstraThreeNodes(t *testing.T) {
	graph := WeightedGraph{map[string][]Edge{}, []string{"a", "b", "c"}}
	source := Edge{vertex: "a"}
	graph.insert(source.vertex, []Edge{{"b", 8, ""}})
	graph.insert("b", []Edge{{"c", 5, ""}})
	actual := graph.dijkstraShortestPath(source)
	expected := []Path{{"a", 0, ""}, {"b", 8, "a"}, {"c", 13, "b"}}
	if !equal(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestDijkstraFourNodesNodes(t *testing.T) {
	graph := WeightedGraph{map[string][]Edge{}, []string{"a", "b", "c", "d"}}
	source := Edge{vertex: "a"}
	graph.insert(source.vertex, []Edge{{"b", 2, ""}, {"c", 4, ""}, {"d", 1, ""}})
	graph.insert("b", []Edge{{"a", 2, ""}, {"c", 1, ""}})
	graph.insert("c", []Edge{{"a", 4, ""}, {"b", 1, ""}})
	graph.insert("d", []Edge{{"a", 1, ""}})
	actual := graph.dijkstraShortestPath(source)
	expected := []Path{{"a", 0, ""}, {"d", 1, "a"}, {"b", 2, "a"}, {"c", 3, "b"}}
	if !equal(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphDijkstraMultipleVertices(t *testing.T) {
	graph := WeightedGraph{map[string][]Edge{}, []string{"a", "b", "c", "d", "e", "f", "g"}}
	source := Edge{vertex: "a"}
	graph.insert("a", []Edge{{"b", 31, ""}, {"c", 2, ""}})
	graph.insert("b", []Edge{{"a", 6, ""}, {"c", 7, ""}, {"d", 1, ""}})
	graph.insert("c", []Edge{{"a", 2, ""}, {"b", 4, ""}, {"d", 5, ""}})
	graph.insert("d", []Edge{{"b", 3, ""}, {"c", 4, ""}, {"e", 8, ""}, {"f", 4, ""}})
	graph.insert("e", []Edge{{"d", 4, ""}, {"g", 3, ""}})
	graph.insert("f", []Edge{{"d", 2, ""}})
	graph.insert("g", []Edge{{"e", 10, ""}})
	var actual = graph.dijkstraShortestPath(source)
	expected := []Path{{"f", 11, "d"}, {"g", 18, "e"}, {"a", 0, ""}, {"b", 6, "c"}, {"c", 2, "a"}, {"d", 7, "c"}, {"e", 15, "d"}}
	if !equal(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphBellmenEmptyGraph(t *testing.T) {
	graph := WeightedGraph{map[string][]Edge{}, []string{}}
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
	graph := WeightedGraph{map[string][]Edge{}, []string{"a", "b"}}
	source := "a"
	graph.insert("a", []Edge{{"b", -1, "a"}})
	var actual, _ = graph.bellmenFord(source)
	expected := []Path{{"a", 0, ""}, {"b", -1, "a"}}
	if !equal(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphEightEdgesFiveVerticesNegativeWeights(t *testing.T) {
	graph := WeightedGraph{map[string][]Edge{}, []string{"a", "b", "c", "d", "e"}}
	source := "a"
	graph.insert("a", []Edge{{"b", -1, "a"}, {"c", 4, "a"}})
	graph.insert("b", []Edge{{"c", 3, "b"}, {"d", 2, "b"}, {"e", 2, "b"}})
	graph.insert("d", []Edge{{"c", 5, "d"}, {"b", 1, "d"}})
	graph.insert("e", []Edge{{"d", -3, "e"}})
	var actual, _ = graph.bellmenFord(source)
	expected := []Path{{"a", 0, ""}, {"b", -1, "a"}, {"c", 2, "b"}, {"d", -2, "e"}, {"e", 1, "b"}}
	if !equal(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestGraphNegativeCycles(t *testing.T) {
	graph := WeightedGraph{map[string][]Edge{}, []string{"a", "b", "c", "d", "e"}}
	source := "a"
	graph.insert("a", []Edge{{"b", 1, "a"}})
	graph.insert("b", []Edge{{"c", 1, "b"}})
	graph.insert("c", []Edge{{"d", -2, "c"}, {"e", 3, "c"}})
	graph.insert("d", []Edge{{"b", -1, "d"}})
	var _, err = graph.bellmenFord(source)
	if err.Error() != "-ve cycle found" {
		t.Errorf("Expected to result in error")
	}
}

func TestFloydEmptyGraph(t *testing.T) {
	graph := WeightedGraph{map[string][]Edge{}, []string{}}
	res := graph.floydWarshall()
	if len(res) != 0 {
		t.Errorf("Expected empty result")
	}
}

func TestFloydTwoVerticesOneEdge(t *testing.T) {
	graph := WeightedGraph{map[string][]Edge{}, []string{"a", "b"}}
	graph.insert("a", []Edge{{"b", 8, "a"}})
	actual := graph.floydWarshall()
	expected := [][]int{{0, 8}, {math.MaxInt32, 0}}
	if !equalMultiDim(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
	}
}

func TestFloydThreeVerticesTwoEdges(t *testing.T) {
	graph := WeightedGraph{map[string][]Edge{}, []string{"a", "b", "c"}}
	graph.insert("a", []Edge{{"b", 8, "a"}, {"c", 1, "a"}})
	graph.insert("b", []Edge{{"c", 4, "b"}})
	actual := graph.floydWarshall()
	expected := [][]int{{0, 8, 1}, {math.MaxInt32, 0, 4}, {math.MaxInt32, math.MaxInt32, 0}}
	if !equalMultiDim(actual, expected) {
		t.Errorf("Expected %v but found %v", expected, actual)
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

func equalMultiDim(actual [][]int, expected [][]int) bool {
	for i := 0; i < len(actual); i++ {
		for j := 0; j < len(actual); j++ {
			if actual[i][j] != expected[i][j] {
				return false
			}
		}
	}
	return true
}
