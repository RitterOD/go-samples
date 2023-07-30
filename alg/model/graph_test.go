package model

import (
	"fmt"
	"testing"
)

func TestGraphCreation(t *testing.T) {

	g := make(GraphRep)
	g[0] = []int{1, 2}
	g[1] = []int{0, 3}
	g[2] = []int{0, 3, 4}
	g[3] = []int{1, 2}
	g[4] = []int{2}
}

func TestGetDotGraphStringRepresentation(t *testing.T) {
	g := NewGraph(DIRECTED)
	g.AddVertex(0, "A")
	g.AddVertex(1, "B")
	g.AddVertex(2, "C")
	g.AddVertex(3, "D")
	g.AddVertex(4, "E")

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 2)

	dotRep := g.getDotRepresentation()
	fmt.Println(dotRep)
}

func TestBFS(t *testing.T) {
	g := NewGraph(DIRECTED)
	g.AddVertex(0, "A")
	g.AddVertex(1, "B")
	g.AddVertex(2, "C")
	g.AddVertex(3, "D")
	g.AddVertex(4, "E")

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 2)

	bfsRes := g.BFS(0)
	dotRep := bfsRes.Graph.getDotRepresentation()
	if bfsRes.Destination[0] != 0 {
		t.Error("Times to 0 !=0")
		return
	}
	if bfsRes.Destination[1] != 1 {
		t.Error("Times to 1 !=1")
		return
	}
	if bfsRes.Destination[2] != 1 {
		t.Error("Times to 2 !=1")
		return
	}

	if bfsRes.Destination[3] != 2 {
		t.Error("Times to 3 !=1")
		return
	}

	if bfsRes.Destination[4] != 2 {
		t.Error("Times to 4 !=2")
		return
	}
	fmt.Println(dotRep)
	weighted := bfsRes.ConvertToWeighted(g)
	fmt.Println(weighted.GetDotRepresentation())
}

func TestUNDIRECTED_GetDotGraphStringRepresentation(t *testing.T) {
	g := NewGraph(UNDIRECTED)
	g.AddVertex(0, "A")
	g.AddVertex(1, "B")
	g.AddVertex(2, "C")
	g.AddVertex(3, "D")
	g.AddVertex(4, "E")

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 2)

	dotRep := g.getDotRepresentation()
	fmt.Println(dotRep)
}

func TestUNDIRECTED_DFS(t *testing.T) {
	g := NewGraph(UNDIRECTED)
	g.AddVertex(0, "A")
	g.AddVertex(1, "B")
	g.AddVertex(2, "C")
	g.AddVertex(3, "D")
	g.AddVertex(4, "E")

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 2)

	dfsResult := g.DFS(0)
	dotRep := dfsResult.Graph.getDotRepresentation()
	fmt.Println(dotRep)
}
