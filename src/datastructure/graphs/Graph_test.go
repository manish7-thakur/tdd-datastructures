package graphs

import (
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
	var actual [0]string
	copy(actual[:], elems)
	if actual != [0]string{} {
		t.Errorf("Expected empty but found %v", elems)
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
