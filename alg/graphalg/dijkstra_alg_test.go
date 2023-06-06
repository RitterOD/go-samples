package graphalg

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
	"sample/alg/dotgraphparser"
	"testing"
)

func TestDijkstra(t *testing.T) {
	//t.Error("TEST_ERROR")
	input, error := antlr.NewFileStream("dijkstra_alg.dot")
	if error != nil {
		t.Error("ERROR OPEN FILE " + error.Error())
		path, err := os.Getwd()
		if err != nil {
			t.Error(err)
		}
		t.Error(path)
		return
	}
	result := dotgraphparser.ParseInputStream(input)
	graph := result.GetGraph()
	if graph.GetName() != "d" {
		t.Error("Error parse weighted graph" + "unexpected name " + graph.GetName())
	} else {
		t.Log("GRAPH NAME PASSED")
	}
	Dijkstra(*graph, "A")

}
