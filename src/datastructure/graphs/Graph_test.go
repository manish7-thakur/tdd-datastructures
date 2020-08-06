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
	if len(elems) != 0 {
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
		t.Errorf("Expected empty but found %v", elems)
	}
}
