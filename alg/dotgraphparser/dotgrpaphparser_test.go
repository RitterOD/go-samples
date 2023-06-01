package dotgraphparser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
	"testing"
)

func TestSimpleWeightGraphParser(t *testing.T) {
	//t.Error("TEST_ERROR")
	input, error := antlr.NewFileStream("sample.dot")
	if error != nil {
		t.Error("ERROR OPEN FILE " + error.Error())
		path, err := os.Getwd()
		if err != nil {
			t.Error(err)
		}
		t.Error(path)
		return
	}
	result := ParseInputStream(input)
	graph := result.graph
	if graph.GetName() != "g" {
		t.Error("Error parse weighted graph" + "unexpected name " + graph.GetName())
	} else {
		t.Log("GRAPH NAME PASSED")
	}
}
