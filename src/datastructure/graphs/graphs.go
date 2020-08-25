package graphs

import (
	"fmt"
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

type Edge struct {
	vertex string
	weight int
	source string
}

type WeightedGraph struct {
	adjList  map[string][]Edge
	vertices []string
}

func (g WeightedGraph) insert(node string, adjacent []Edge) {
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

func (g WeightedGraph) dijkstraShortestPath(source Edge) []Path {
	var minQueue []Path
	var processed []Path
	pathList := make(map[string]*Path)
	for _, v := range g.vertices {
		if v == source.vertex {
			pathList[v] = &Path{v, 0, ""}
			minQueue = append(minQueue, *pathList[v])
		} else {
			pathList[v] = &Path{v, math.MaxInt32, ""}
		}
	}
	var visited = make(map[string]struct{})
	for len(minQueue) != 0 {
		current := minQueue[0]
		minQueue = minQueue[1:]
		src := current.vertex
		adjacent := g.adjList[src]
		for _, n := range adjacent {
			dest := n.vertex
			weight := n.weight
			path := pathList[dest]
			if _, ok := visited[dest]; !ok { // to optimize a bit
				if pathList[src].dist+weight < path.dist {
					path.dist = pathList[src].dist + weight
					path.predecessor = src
					minQueue = appendMin(minQueue, *path)
				}
			}
		}
		visited[src] = struct{}{}
	}
	for _, path := range pathList {
		processed = append(processed, *path)
	}
	return processed
}

type Path struct {
	vertex      string
	dist        int
	predecessor string
}

func (g WeightedGraph) bellmenFord(source string) ([]Path, error) {
	pathList := make(map[string]*Path)
	for _, v := range g.vertices {
		if v == source {
			pathList[v] = &Path{v, 0, ""}
		} else {
			pathList[v] = &Path{v, math.MaxInt32, ""}
		}
	}
	for i := 0; i < len(g.vertices); i++ {
		for _, edges := range g.adjList {
			for _, edge := range edges {
				src := edge.source
				dest := edge.vertex
				weight := edge.weight
				path := pathList[dest]
				if pathList[src].dist+weight < path.dist {
					path.dist = pathList[src].dist + weight
					path.predecessor = src
				}
			}
		}
	}
	for _, edges := range g.adjList {
		for _, edge := range edges {
			src := edge.source
			dest := edge.vertex
			weight := edge.weight
			path := pathList[dest]
			if pathList[src].dist+weight < path.dist {
				return []Path{}, fmt.Errorf("-ve cycle found")
			}
		}
	}
	var processed []Path
	for _, path := range pathList {
		processed = append(processed, *path)
	}
	return processed, nil
}

func appendMin(queue []Path, adjacent ...Path) []Path {
	queue = append(queue, adjacent...)
	sort.Slice(queue, func(i, j int) bool {
		return queue[i].dist < queue[j].dist
	})
	return queue
}

/* Can be done using min heap, reference : https://golang.org/pkg/container/heap/#example__intHeap
type NodeHeap []Edge

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
	*h = append(*h, x.(Edge))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
*/
