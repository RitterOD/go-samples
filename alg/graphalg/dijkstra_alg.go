package graphalg

import (
	"fmt"
	pq "github.com/jupp0r/go-priority-queue"
	"math"
	"sample/alg/model"
)

func Dijkstra(graph model.WeightedGraph, srcVertexName string) model.ShortestPathResult {

	rv := make(model.ShortestPathResult)
	rep := graph.GetRepresentation()
	for v, _ := range rep {
		rv[v] = &model.ShortestPathNode{Vertex: v, Parent: 0, Destination: math.NaN()}
	}
	srcV := graph.GetVertexIndexByName(srcVertexName)
	rv[srcV].Destination = 0
	//var rv model.ShortestPathResult
	var vertexPQ = pq.New()
	for _, r := range rv {
		vertexPQ.Insert(r, r.Destination)
	}
	for vertexPQ.Len() != 0 {
		tmp, _ := vertexPQ.Pop()
		node, _ := tmp.(model.ShortestPathNode)
		edges := rep[node.Vertex]
		fmt.Println(edges)
	}
	//graph.
	return rv
}

func relax(u, v int, weight float64) {

}
