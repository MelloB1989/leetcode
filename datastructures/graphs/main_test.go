package main

import (
	"testing"

	graph "graph/graph_implementation"
)

func TestAddVertex(t *testing.T) {
	g := graph.NewGraph()
	g.AddVertex("A")

	if !g.HasVertex("A") {
		t.Errorf("Expected vertex A to be present")
	}
}

func TestAddEdge(t *testing.T) {
	g := graph.NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B")

	if !g.HasEdge("A", "B") {
		t.Errorf("Expected edge A -> B to be present")
	}
}

func TestRemoveVertex(t *testing.T) {
	g := graph.NewGraph()
	g.AddVertex("A")
	g.RemoveVertex("A")

	if g.HasVertex("A") {
		t.Errorf("Expected vertex A to be removed")
	}
}

func TestRemoveEdge(t *testing.T) {
	g := graph.NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B")
	g.RemoveEdge("A", "B")

	if g.HasEdge("A", "B") {
		t.Errorf("Expected edge A -> B to be removed")
	}
}

func TestHasVertex(t *testing.T) {
	g := graph.NewGraph()
	g.AddVertex("X")

	if !g.HasVertex("X") {
		t.Errorf("Expected vertex X to exist")
	}

	if g.HasVertex("Y") {
		t.Errorf("Did not expect vertex Y to exist")
	}
}

func TestHasEdge(t *testing.T) {
	g := graph.NewGraph()
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
	g := graph.NewGraph()
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
	g := graph.NewGraph()
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
	g := graph.NewGraph()
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
	g := graph.NewGraph()
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
	g := graph.NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")

	g.AddEdge("A", "B")
	g.AddEdge("B", "C")
	g.AddEdge("C", "A") // creates a cycle

	if !g.IsCyclic() {
		t.Errorf("Expected cycle to be detected")
	}

	g2 := graph.NewGraph()
	g2.AddVertex("X")
	g2.AddVertex("Y")
	g2.AddVertex("Z")

	g2.AddEdge("X", "Y")
	g2.AddEdge("Y", "Z")

	if g2.IsCyclic() {
		t.Errorf("Did not expect a cycle")
	}
}
