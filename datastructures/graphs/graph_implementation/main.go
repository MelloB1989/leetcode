package graph_implementation

import "slices"

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

func (g *Graph) IsCyclic() bool {
	return false
}
