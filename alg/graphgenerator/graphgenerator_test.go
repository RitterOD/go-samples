package graphgenerator

import (
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {
	//t.Error("TEST_ERROR")
	graph := ClusteredRandomGraph(5, 5, 1, 1, 0.5, 0.5, 1, 100)
	fmt.Println(graph.GetDotRepresentation())

}
