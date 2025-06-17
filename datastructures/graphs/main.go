package graph

import (
	"slices"

	"datastructures/queue"
)

type Graph struct {
	adjacencyList map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[string][]string),
	}
}

func (g *Graph) AddVertex(vertex string) {
	g.adjacencyList[vertex] = []string{}
}

func (g *Graph) AddEdge(from, to string) {
	g.adjacencyList[from] = append(g.adjacencyList[from], to)
}

func (g *Graph) RemoveVertex(vertex string) {
	for v, _ := range g.adjacencyList {
		g.RemoveEdge(v, vertex)
	}
	delete(g.adjacencyList, vertex)
}

func (g *Graph) RemoveEdge(from, to string) {
	res := []string{}
	for _, edge := range g.adjacencyList[from] {
		if edge == to {
			continue
		}
		res = append(res, edge)
	}
	g.adjacencyList[from] = res
}

func (g *Graph) HasVertex(vertex string) bool {
	_, ok := g.adjacencyList[vertex]
	return ok
}

func (g *Graph) HasEdge(from, to string) bool {
	if slices.Contains(g.adjacencyList[from], to) {
		return true
	}
	return false
}

func (g *Graph) GetAdjacencyList(vertex string) []string {
	return g.adjacencyList[vertex]
}

func (g *Graph) Vertices() []string {
	vertices := []string{}
	for k, _ := range g.adjacencyList {
		vertices = append(vertices, k)
	}
	return vertices
}

func (g *Graph) Edges() [][2]string {
	edges := [][2]string{}
	for from, v := range g.adjacencyList {
		for _, to := range v {
			edges = append(edges, [2]string{from, to})
		}
	}
	return edges
}

func (g *Graph) InDegree(vertex string) int {
	inDegree := 0
	for _, edges := range g.adjacencyList {
		if slices.Contains(edges, vertex) {
			inDegree++
		}
	}
	return inDegree
}

func (g *Graph) OutDegree(vertex string) int {
	return len(g.adjacencyList[vertex])
}

func (g *Graph) BFS(start string) []string {
	queue := queue.NewQueue()
	res := []string{}
	visited := map[string]bool{}
	queue.Push(start)
	visited[start] = true
	for !queue.IsEmpty() {
		v := queue.Poll()
		res = append(res, v.(string))
		for _, u := range g.adjacencyList[v.(string)] {
			if !visited[u] {
				queue.Push(u)
				visited[u] = true
			}
		}
	}
	return res
}
func (g *Graph) DFS(start string) []string {
	visited := map[string]bool{}
	res := []string{}
	var dfsr func(v string)
	dfsr = func(v string) {
		visited[v] = true
		res = append(res, v)
		for _, u := range g.adjacencyList[v] {
			if !visited[u] {
				dfsr(u)
			}
		}
	}
	dfsr(start)
	return res
}

// Undirected
//
//	func (g *Graph) IsCyclic() bool {
//		visited := map[string]bool{}
//		var dfs func(parent, s string) bool
//		dfs = func(parent string, s string) bool {
//			visited[s] = true
//			for _, u := range g.adjacencyList[s] {
//				if !visited[u] {
//					if dfs(s, u) {
//						return true
//					}
//				} else if parent != u {
//					return true
//				}
//			}
//			return false
//		}
//		for k, _ := range g.adjacencyList {
//			if !visited[k] {
//				if dfs("", k) {
//					return true
//				}
//			}
//		}
//		return false
//	}
//
// Directed
func (g *Graph) IsCyclic() bool {
	visited := map[string]bool{}
	recStack := map[string]bool{}
	var dfs func(s string) bool
	dfs = func(s string) bool {
		visited[s] = true
		recStack[s] = true
		for _, u := range g.adjacencyList[s] {
			if !visited[u] && dfs(u) {
				return true
			} else if recStack[u] {
				return true
			}
		}
		recStack[s] = false
		return false
	}
	for v, _ := range g.adjacencyList {
		if !visited[v] && dfs(v) {
			return true
		}
	}
	return false
}

func (g *Graph) ShortestPath(from, to string) []string {
	visited := map[string]bool{}
	prev := map[string]string{}
	queue := queue.NewQueue()
	visited[from] = true
	queue.Push(from)
	for !queue.IsEmpty() {
		curr := queue.Poll().(string)
		for _, neighbour := range g.adjacencyList[curr] {
			prev[neighbour] = curr
			visited[neighbour] = true
			queue.Push(neighbour)
			if neighbour == to {
				queue.Flush()
				break
			}
		}
	}
	path := []string{}
	at := to
	for {
		path = append([]string{at}, path...)
		if at == from {
			break
		}
		if p, ok := prev[at]; ok {
			at = p
		} else {
			return nil
		}
	}
	return path
}

func (g *Graph) ShortestPathFromAVertex(from string) map[string]int {
	dist := make(map[string]int)
	for v := range g.adjacencyList {
		dist[v] = -1
	}
	queue := queue.NewQueue()
	dist[from] = 0
	queue.Push(from)

	for !queue.IsEmpty() {
		curr := queue.Poll().(string)
		for _, neighbor := range g.adjacencyList[curr] {
			if dist[neighbor] == -1 {
				dist[neighbor] = dist[curr] + 1
				queue.Push(neighbor)
			}
		}
	}

	return dist
}
