package model

import (
	"container/list"
	"fmt"
	"strings"
)

type GraphRep map[int][]int

type GraphType int64

type Edge struct {
	u int
	v int
}

const (
	DIRECTED          GraphType = 0
	UNDIRECTED        GraphType = 1
	WEIGHTED_DIRECTED           = 2
)

type DfsColor int64

const (
	WHITE DfsColor = 0
	GREY  DfsColor = 1
	BLACK DfsColor = 2
)

type DfsTime struct {
	tIn  int64
	tOut int64
}

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

type DFSresult struct {
	graph  Graph
	source int
	times  map[int]DfsTime
}

type WeighedEdge struct {
	VertexIndex int
	Weight      float64
}

type WeightedGraphRep map[int][]WeighedEdge

type WeightedGraph struct {
	rep          WeightedGraphRep
	name         string
	vertexToName map[int]string
	nameToVertex map[string]int
	graphType    GraphType
}

func NewWeightedGraph() *WeightedGraph {
	g := new(WeightedGraph)
	g.rep = make(WeightedGraphRep)
	g.graphType = WEIGHTED_DIRECTED
	g.nameToVertex = make(map[string]int)
	g.vertexToName = make(map[int]string)
	return g
}

func (g *WeightedGraph) Name(name string) {
	g.name = name
}
func (g *WeightedGraph) GetName() string {
	return g.name
}

func (g *WeightedGraph) AddVertex(v int, name string) {
	_, existed := g.rep[v]
	if !existed {
		g.rep[v] = []WeighedEdge{}
		g.nameToVertex[name] = v
		g.vertexToName[v] = name
	}
}

func (g *WeightedGraph) AddEdge(u, v int, weight float64) {
	g.rep[v] = append(g.rep[v], WeighedEdge{VertexIndex: u, Weight: weight})
}

func (g *WeightedGraph) AddEdgeByVertexName(uName, vName string, weight float64) {
	u := g.nameToVertex[uName]
	v := g.nameToVertex[vName]
	g.rep[u] = append(g.rep[u], WeighedEdge{VertexIndex: v, Weight: weight})
}

func (g *WeightedGraph) GetDotRepresentation() string {
	nodes := make([]string, 0)
	for n, _ := range g.nameToVertex {
		nodes = append(nodes, n)
	}
	edges := make([]string, 0)
	for v, adjNodes := range g.rep {
		for _, u := range adjNodes {
			edges = append(edges, g.vertexToName[v]+" -> "+g.vertexToName[u.VertexIndex])
		}
	}
	return "digraph" + g.name + " { " + strings.Join(nodes[:], ";\n") + "\n" + strings.Join(edges[:], "\n") + " }"
}

func (g *WeightedGraph) GetRepresentation() WeightedGraphRep {
	return g.rep
}

func (g *WeightedGraph) GetVertexIndexByName(name string) int {
	return g.nameToVertex[name]
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
	_, existed := g.rep[v]
	if !existed {
		g.rep[v] = []int{}
		g.nameToVertex[name] = v
		g.vertexToName[v] = name
	}
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
	} else if g.graphType == UNDIRECTED {
		edgeSet := make(map[Edge]string)
		for v, adjNodes := range g.rep {
			for _, u := range adjNodes {
				var edge Edge
				if v < u {
					edge = Edge{v, u}
				} else {
					edge = Edge{u, v}
				}
				edgeSet[edge] = g.vertexToName[u] + " -- " + g.vertexToName[v]
			}
		}
		for _, e := range edgeSet {
			edges = append(edges, e)
		}
		return "graph { " + strings.Join(nodes[:], ";\n") + "\n" + strings.Join(edges[:], "\n") + " }"
	} else {
		panic("UNSUPPORTED GRAPH TYPE")
	}
}

func (g Graph) BFS(s int) BFSresult {
	result := new(BFSresult)
	result.graph = *NewGraph(DIRECTED)
	result.destination = make(map[int]int)
	visitNodes := make([]int, len(g.rep))
	rg := result.graph
	queue := list.New()
	queue.PushBack(s)
	result.destination[s] = 0
	visitNodes[s] = 1
	for queue.Len() != 0 {
		cur := queue.Front()
		rg.AddVertex(cur.Value.(int), g.vertexToName[cur.Value.(int)])
		for _, u := range g.rep[cur.Value.(int)] {
			if visitNodes[u] == 0 {
				visitNodes[u] = 1
				rg.AddEdge(cur.Value.(int), u)
				queue.PushBack(u)
				result.destination[u] = result.destination[cur.Value.(int)] + 1
			}
		}
		queue.Remove(cur)
	}
	return *result
}

func (g Graph) DFS(s int) DFSresult {
	result := new(DFSresult)
	result.graph = *NewGraph(DIRECTED)
	result.times = make(map[int]DfsTime)
	visitNodes := make([]DfsColor, len(g.rep))
	rg := result.graph
	stack := list.New()
	stack.PushFront(s)
	time := int64(0)
	result.times[s] = DfsTime{time, -1}
	visitNodes[s] = GREY
	for stack.Len() != 0 {
		cur := stack.Front()
		rg.AddVertex(cur.Value.(int), g.vertexToName[cur.Value.(int)])
		allVisited := true
		for _, u := range g.rep[cur.Value.(int)] {
			if visitNodes[u] == WHITE {
				time++
				visitNodes[u] = GREY
				rg.AddEdge(cur.Value.(int), u)
				stack.PushFront(u)
				result.times[u] = DfsTime{time, -1}
				allVisited = false
				break
			}
		}
		if allVisited {
			stack.Remove(cur)
			time++
			tmpTime := result.times[cur.Value.(int)]
			tmpTime.tOut = time
			visitNodes[cur.Value.(int)] = BLACK
			result.times[cur.Value.(int)] = tmpTime
		}
	}
	return *result
}

type ShortestPathNode struct {
	Vertex      int
	Parent      int
	Marked      bool
	Destination float64
}

type ShortestPathResult map[int]*ShortestPathNode

func (result *ShortestPathResult) ConvertToWeighted() *WeightedGraph {
	rv := NewWeightedGraph()
	for v, r := range *result {
		rv.AddVertex(v, fmt.Sprint(v))
		rv.AddVertex(r.Parent, fmt.Sprint(r.Parent))
		rv.AddEdge(r.Parent, v, r.Destination-(*result)[r.Parent].Destination)
	}
	return rv
}
