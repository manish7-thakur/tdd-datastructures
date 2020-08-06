package graphs

type Graph struct {
	adjList map[string][]string
}

func (graph Graph) insert(u string, v string) {
	graph.adjList[u] = append(graph.adjList[u], v)
}

func (graph Graph) bfs(root string) []string {
	var queue []string
	var visited []string
	queue = append(queue, root)
	for ; len(queue) != 0; {
		current := queue[0]
		visited = append(visited, current)
		adjacent := graph.adjList[current]
		queue = append(queue, adjacent...)
		queue = queue[1:]
	}
	return visited
}
