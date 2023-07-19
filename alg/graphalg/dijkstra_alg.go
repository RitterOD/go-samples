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
		tmpNode, ok := tmp.(*model.ShortestPathNode)
		if !ok {
			log.Fatal("Can't convert")
		} else {
			for v, _ := range rep[tmpNode.Vertex] {
				relax(v, tmpNode.Vertex, 0, &rv)
			}
		}
		fmt.Println(tmp)
	}
	//graph.
	return rv
}

func relax(u, v int, weight float64, shortestPath *model.ShortestPathResult) {
	if (*shortestPath)[u].Destination > (*shortestPath)[v].Destination+weight {
		(*shortestPath)[u] = &model.ShortestPathNode{Parent: v, Destination: (*shortestPath)[v].Destination + weight}
	}

}
