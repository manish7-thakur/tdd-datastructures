package graphs

type Graph struct {
	adjList map[string][]string
}

func (graph Graph) insert(u string, v string) {
	graph.adjList[u] = append(graph.adjList[u], v)
}

func (graph Graph) bfs(root string) []string {
	strings := graph.adjList[root]
	return strings
}