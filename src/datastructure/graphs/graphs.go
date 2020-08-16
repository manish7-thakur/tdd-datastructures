package graphs

import (
	"math"
	"sort"
)

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

func (g WeightedGraph) bellmenFord(source Node) []Node {
	source.dist = 0
	var edges []Node
	v := len(g.adjList)
	for _, edgeList := range g.adjList {
		for _, e := range edgeList {
			if e.vertex == source.vertex {
				e.dist = 0
			} else {
				e.dist = math.MaxInt32
			}
			edges = append(edges, e)
		}
	}
	for i := 0; i <= v; i = i + 1 {
		for _, edge := range edges {
			if source.dist + edge.weight < edge.dist {
				edges[i].dist = source.dist + edge.weight
			}
		}
	}
	return edges
}

func appendMin(queue []Node, adjacent ...Node) []Node {
	queue = append(queue, adjacent...)
	sort.Slice(queue, func(i, j int) bool {
		return queue[i].dist < queue[j].dist
	})
	return queue
}

/* Can be done using min heap, reference : https://golang.org/pkg/container/heap/#example__intHeap
type NodeHeap []Node

func (h NodeHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h NodeHeap) Len() int {
	return len(h)
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
*/
