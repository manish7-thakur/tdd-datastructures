package graphs

type Graph struct {
	adjList map[string][]string
}

func (graph Graph) insert(u string, v string) {
	graph.adjList[u] = append(graph.adjList[u], v)
}

func (graph Graph) insertAdj(u string, adjacent []string) {
	graph.adjList[u] = append(graph.adjList[u], adjacent...)
}

func (graph Graph) bfs(root string) []string {
	var queue []string
	var processed []string
	visited := make(map[string]struct{})
	queue = append(queue, root)
	for ; len(queue) != 0; {
		current := queue[0]
		_, ok := visited[current]
		if !ok {
			visited[current] = struct{}{}
			processed = append(processed, current)
			adjacent := graph.adjList[current]
			queue = append(queue, adjacent...)
		}

		queue = queue[1:]
	}
	return processed
}
