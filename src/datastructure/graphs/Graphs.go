package graphs

import "sort"

type Graph struct {
	adjList map[string][]string
}

func (graph Graph) insert(u string, v string) {
	graph.adjList[u] = append(graph.adjList[u], v)
}

func (graph Graph) insertAdj(u string, adjacent []string) {
	graph.adjList[u] = append(graph.adjList[u], adjacent...)
}

type Node struct {
	vertex string
	weight int
	dist   int
}

type WeightedGraph struct {
	adjList map[string][]Node
}

func (g WeightedGraph) insert(node string, adjacent []Node) {
	g.adjList[node] = adjacent
}

func (graph Graph) bfs(start string) []string {
	var queue []string
	var processed []string
	visited := make(map[string]struct{})
	queue = append(queue, start)
	for ; len(queue) != 0; {
		current := queue[0]
		queue = queue[1:]
		_, ok := visited[current]
		if !ok {
			visited[current] = struct{}{}
			processed = append(processed, current)
			queue = append(queue, graph.adjList[current]...)
		}
	}
	return processed
}

func (graph Graph) dfs(start string) []string {
	var stack []string
	var processed []string
	visited := make(map[string]struct{})
	stack = append(stack, start)
	for ; len(stack) != 0; {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		_, ok := visited[current]
		if !ok {
			visited[current] = struct{}{}
			processed = append(processed, current)
			stack = append(stack, graph.adjList[current]...)
		}
	}
	return processed
}

func (g WeightedGraph) dijkstraShortestPath(source Node) []Node {
	var minQueue []Node
	var processed []Node
	var visited = make(map[string]struct{})
	source.dist = 0
	minQueue = append(minQueue, source)
	for len(minQueue) != 0 {
		current := minQueue[0]
		minQueue = minQueue[1:]
		adjacent := g.adjList[current.vertex]
		for i, n := range adjacent {
			if _, ok := visited[n.vertex]; !ok {
				if current.dist+n.weight < n.dist {
					adjacent[i].dist = current.dist + n.weight
					minQueue = appendMin(minQueue, adjacent[i])
				}
			}
		}
		if _, ok := visited[current.vertex]; !ok {
			processed = append(processed, current)
		}
		visited[current.vertex] = struct{}{}
	}
	return processed
}

func appendMin(queue []Node, adjacent ...Node) []Node {
	queue = append(queue, adjacent...)
	sort.Slice(queue, func(i, j int) bool {
		return queue[i].dist < queue[j].dist
	})
	return queue
}
