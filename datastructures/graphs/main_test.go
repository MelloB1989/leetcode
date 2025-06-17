package graph

import (
	"reflect"
	"testing"
)

func TestAddVertex(t *testing.T) {
	g := NewGraph()
	g.AddVertex("A")

	if !g.HasVertex("A") {
		t.Errorf("Expected vertex A to be present")
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B")

	if !g.HasEdge("A", "B") {
		t.Errorf("Expected edge A -> B to be present")
	}
}

func TestRemoveVertex(t *testing.T) {
	g := NewGraph()
	g.AddVertex("A")
	g.RemoveVertex("A")

	if g.HasVertex("A") {
		t.Errorf("Expected vertex A to be removed")
	}
}

func TestRemoveEdge(t *testing.T) {
	g := NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B")
	g.RemoveEdge("A", "B")

	if g.HasEdge("A", "B") {
		t.Errorf("Expected edge A -> B to be removed")
	}
}

func TestHasVertex(t *testing.T) {
	g := NewGraph()
	g.AddVertex("X")

	if !g.HasVertex("X") {
		t.Errorf("Expected vertex X to exist")
	}

	if g.HasVertex("Y") {
		t.Errorf("Did not expect vertex Y to exist")
	}
}

func TestHasEdge(t *testing.T) {
	g := NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B")

	if !g.HasEdge("A", "B") {
		t.Errorf("Expected edge A -> B to exist")
	}

	if g.HasEdge("B", "A") {
		t.Errorf("Did not expect edge B -> A to exist (Directed Graph)")
	}
}

func TestGetAdjacencyList(t *testing.T) {
	g := NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")

	adj := g.GetAdjacencyList("A")
	expected := map[string]bool{"B": true, "C": true}

	for _, v := range adj {
		if !expected[v] {
			t.Errorf("Unexpected adjacent vertex: %s", v)
		}
	}
}

func TestVertices(t *testing.T) {
	g := NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")

	vertices := g.Vertices()
	expected := map[string]bool{"A": true, "B": true}

	if len(vertices) != len(expected) {
		t.Errorf("Expected %d vertices, got %d", len(expected), len(vertices))
	}
	for _, v := range vertices {
		if !expected[v] {
			t.Errorf("Unexpected vertex: %s", v)
		}
	}
}

func TestEdges(t *testing.T) {
	g := NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B")

	edges := g.Edges()
	expected := [][2]string{{"A", "B"}}

	if len(edges) != len(expected) {
		t.Errorf("Expected %d edges, got %d", len(expected), len(edges))
	}
	found := false
	for _, e := range edges {
		if e == expected[0] {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected edge (A, B) not found")
	}
}

func TestInOutDegree(t *testing.T) {
	g := NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddEdge("A", "B")
	g.AddEdge("C", "B")

	if g.InDegree("B") != 2 {
		t.Errorf("Expected in-degree of B to be 2")
	}
	if g.OutDegree("A") != 1 {
		t.Errorf("Expected out-degree of A to be 1")
	}
	if g.OutDegree("B") != 0 {
		t.Errorf("Expected out-degree of B to be 0")
	}
}

func TestIsCyclic(t *testing.T) {
	g := NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")

	g.AddEdge("A", "B")
	g.AddEdge("B", "C")
	g.AddEdge("C", "A") // creates a cycle

	if !g.IsCyclic() {
		t.Errorf("Expected cycle to be detected")
	}

	g2 := NewGraph()
	g2.AddVertex("X")
	g2.AddVertex("Y")
	g2.AddVertex("Z")

	g2.AddEdge("X", "Y")
	g2.AddEdge("Y", "Z")

	if g2.IsCyclic() {
		t.Errorf("Did not expect a cycle")
	}
}

func setupGraph() *Graph {
	g := NewGraph()
	for _, v := range []string{"A", "B", "C", "D", "E"} {
		g.AddVertex(v)
	}
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("C", "D")
	g.AddEdge("A", "D")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")
	return g
}

func TestBFS(t *testing.T) {
	g := setupGraph()
	result := g.BFS("A")
	expected := []string{"A", "B", "C", "D", "E"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BFS failed. Expected %v, got %v", expected, result)
	}
}

func TestDFS(t *testing.T) {
	g := setupGraph()
	result := g.DFS("A")
	// Depending on implementation order, one valid result is:
	expected := []string{"A", "B", "D", "C", "E"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("DFS failed. Expected %v, got %v", expected, result)
	}
}

func TestShortestPath(t *testing.T) {
	g := setupGraph()

	tests := []struct {
		from, to string
		expected []string
	}{
		{"A", "D", []string{"A", "D"}},
		{"A", "E", []string{"A", "C", "E"}},
		{"A", "A", []string{"A"}},
		{"B", "E", nil}, // no path
	}

	for _, tc := range tests {
		result := g.ShortestPath(tc.from, tc.to)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("ShortestPath(%s -> %s) = %v, want %v", tc.from, tc.to, result, tc.expected)
		}
	}
}

func TestShortestPathFromAVertex(t *testing.T) {
	g := NewGraph()
	for _, v := range []string{"A", "B", "C", "D", "E", "F"} {
		g.AddVertex(v)
	}
	// Build the graph:
	// A → B → D
	// A → C → E
	// F is disconnected

	g.AddEdge("A", "B")
	g.AddEdge("B", "D")
	g.AddEdge("A", "C")
	g.AddEdge("C", "E")

	expected := map[string]int{
		"A": 0,
		"B": 1,
		"C": 1,
		"D": 2,
		"E": 2,
		"F": -1, // unreachable
	}

	result := g.ShortestPathFromAVertex("A")

	for node, dist := range expected {
		if result[node] != dist {
			t.Errorf("ShortestPathFromAVertex[A → %s] = %d; want %d", node, result[node], dist)
		}
	}
}
