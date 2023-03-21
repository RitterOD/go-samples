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
