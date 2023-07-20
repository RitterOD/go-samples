package graphalg

import (
	"fmt"
	pq "github.com/jupp0r/go-priority-queue"
	"log"
	"math"
	"sample/alg/model"
)

func Dijkstra(graph model.WeightedGraph, srcVertexName string) model.ShortestPathResult {

	rv := make(model.ShortestPathResult)
	rep := graph.GetRepresentation()
	for v, _ := range rep {
		rv[v] = &model.ShortestPathNode{Vertex: v, Parent: 0, Destination: math.MaxFloat64, Marked: false}
	}
	srcV := graph.GetVertexIndexByName(srcVertexName)
	rv[srcV].Destination = 0
	//var rv model.ShortestPathResult
	var vertexPQ = pq.New()
	vertexPQ.Insert(rv[srcV], 0)
	for vertexPQ.Len() != 0 {
		tmp, _ := vertexPQ.Pop()
		tmpNode, ok := tmp.(*model.ShortestPathNode)
		if !ok {
			log.Fatal("Can't convert")
		} else {
			if !tmpNode.Marked {
				for _, e := range rep[tmpNode.Vertex] {
					isRel := relax(tmpNode.Vertex, e.VertexIndex, e.Weight, &rv)
					if isRel {
						vertexPQ.Insert(rv[e.VertexIndex], rv[e.VertexIndex].Destination)
					}
				}
				tmpNode.Marked = true
			}
		}
		fmt.Println(tmp)
	}
	//graph.
	return rv
}

func relax(u, v int, weight float64, shortestPath *model.ShortestPathResult) bool {
	if (*shortestPath)[v].Destination > (*shortestPath)[u].Destination+weight {
		(*shortestPath)[v] = &model.ShortestPathNode{Vertex: v, Parent: u, Destination: (*shortestPath)[u].Destination + weight, Marked: false}
		return true
	} else {
		return false
	}

}
