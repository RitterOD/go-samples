package model

import (
	"container/list"
	"strings"
)

type GraphRep map[int][]int

type GraphType int64

const (
	DIRECTED   GraphType = 0
	UNDIRECTED GraphType = 1
)

type Graph struct {
	rep          GraphRep
	vertexToName map[int]string
	nameToVertex map[string]int
	graphType    GraphType
}

type BFSresult struct {
	graph       Graph
	source      int
	destination map[int]int
}

func NewGraph(graphType GraphType) *Graph {
	g := new(Graph)
	g.rep = make(GraphRep)
	g.graphType = graphType
	g.nameToVertex = make(map[string]int)
	g.vertexToName = make(map[int]string)
	return g
}

// TODO ERROR CHECK

func (g Graph) AddVertex(v int, name string) {
	g.rep[v] = []int{}
	g.nameToVertex[name] = v
	g.vertexToName[v] = name
}

// TODO ERROR CHECK
func (g Graph) AddEdge(u, v int) {
	if g.graphType == DIRECTED {
		g.rep[u] = append(g.rep[u], v)
	} else if g.graphType == UNDIRECTED {
		g.rep[v] = append(g.rep[v], u)
		g.rep[u] = append(g.rep[u], v)
	} else {
		panic("Not supported graph type")
	}
}

func (g Graph) getDotRepresentation() string {
	nodes := make([]string, 0)
	for n, _ := range g.nameToVertex {
		nodes = append(nodes, n)
	}
	edges := make([]string, 0)
	if g.graphType == DIRECTED {
		for v, adjNodes := range g.rep {
			for _, u := range adjNodes {
				edges = append(edges, g.vertexToName[v]+" -> "+g.vertexToName[u])
			}
		}
		return "digraph { " + strings.Join(nodes[:], ";\n") + "\n" + strings.Join(edges[:], "\n") + " }"
	} else {
		panic("UNSUPPORTED GRAPH TYPE")
	}

}

func (g Graph) BFS(s int) Graph {
	queue := list.New()
	queue.PushBack(s)
	dst := 0
	for queue.Len() != 0 {
		dst++
	}
	panic("Not implemented")
}
