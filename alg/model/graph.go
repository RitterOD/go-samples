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
	TimeIn  int64
	TimeOut int64
}

type Graph struct {
	rep          GraphRep
	vertexToName map[int]string
	nameToVertex map[string]int
	graphType    GraphType
}

type BFSresult struct {
	Graph       Graph
	Source      int
	Destination map[int]int
	Parent      map[int]int
}

type DFSresult struct {
	Graph  Graph
	Source int
	Times  map[int]DfsTime
	Parent map[int]int
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
			edges = append(edges, g.vertexToName[v]+" -> "+g.vertexToName[u.VertexIndex]+
				"[weight ="+fmt.Sprint(u.Weight)+" ]")
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
		panic("Not supported Graph type")
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
		return "Graph { " + strings.Join(nodes[:], ";\n") + "\n" + strings.Join(edges[:], "\n") + " }"
	} else {
		panic("UNSUPPORTED GRAPH TYPE")
	}
}

func (g Graph) BFS(s int) BFSresult {
	result := new(BFSresult)
	result.Graph = *NewGraph(DIRECTED)
	result.Destination = make(map[int]int)
	result.Parent = make(map[int]int)
	visitNodes := make([]int, len(g.rep))
	rg := result.Graph
	queue := list.New()
	queue.PushBack(s)
	result.Destination[s] = 0
	visitNodes[s] = 1
	for queue.Len() != 0 {
		cur := queue.Front()
		rg.AddVertex(cur.Value.(int), g.vertexToName[cur.Value.(int)])
		for _, u := range g.rep[cur.Value.(int)] {
			if visitNodes[u] == 0 {
				visitNodes[u] = 1
				rg.AddEdge(cur.Value.(int), u)
				queue.PushBack(u)
				result.Destination[u] = result.Destination[cur.Value.(int)] + 1
				result.Parent[u] = cur.Value.(int)
			}
		}
		queue.Remove(cur)
	}
	return *result
}

func (g *Graph) DFS(s int) *DFSresult {
	result := new(DFSresult)
	result.Graph = *NewGraph(DIRECTED)
	result.Times = make(map[int]DfsTime)
	visitNodes := make([]DfsColor, len(g.rep))
	rg := result.Graph
	stack := list.New()
	stack.PushFront(s)
	time := int64(0)
	result.Times[s] = DfsTime{time, -1}
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
				result.Times[u] = DfsTime{time, -1}
				allVisited = false
				break
			}
		}
		if allVisited {
			stack.Remove(cur)
			time++
			tmpTime := result.Times[cur.Value.(int)]
			tmpTime.TimeOut = time
			visitNodes[cur.Value.(int)] = BLACK
			result.Times[cur.Value.(int)] = tmpTime
		}
	}
	return result
}

type ShortestPathNode struct {
	Vertex      int
	Parent      int
	Marked      bool
	Destination float64
}

type ShortestPathResult map[int]*ShortestPathNode

func (result *ShortestPathResult) ConvertToWeighted(srcGraph *WeightedGraph) *WeightedGraph {
	rv := NewWeightedGraph()
	for v, r := range *result {
		rv.AddVertex(v, srcGraph.vertexToName[v])
		rv.AddVertex(r.Parent, srcGraph.vertexToName[r.Parent])
		if r.Parent != v {
			rv.AddEdge(v, r.Parent, r.Destination-(*result)[r.Parent].Destination)
		}
	}
	return rv
}

func (result *BFSresult) ConvertToWeighted(srcGraph *Graph) *WeightedGraph {
	rv := NewWeightedGraph()
	for v, p := range result.Parent {
		rv.AddVertex(v, srcGraph.vertexToName[v])
		rv.AddVertex(p, srcGraph.vertexToName[p])
		if v != p {
			rv.AddEdge(p, v, float64(result.Destination[v]-result.Destination[p]))
		}
	}
	return rv
}
