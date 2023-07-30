package graphgenerator

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestGenerator(t *testing.T) {
	graph := ClusteredRandomGraph(5, 5, 1, 1, 0.5, 0.5, 1, 100, 1000, 1500)
	fmt.Println(graph.GetDotRepresentation())
}

func TestGeneratorTwoClusters(t *testing.T) {
	graph := ClusteredRandomGraph(5, 5, 2, 2, 0.5, 1, 1, 100, 1000, 1500)
	fmt.Println(graph.GetDotRepresentation())
}

func TestGenerateToFile(t *testing.T) {
	graph := ClusteredRandomGraph(100, 100, 1, 1, 0.5, 1, 1, 100, 1000, 1500)
	//fmt.Println(graph.GetDotRepresentation())
	fileName := "oneCluster_" + time.Now().String() + ".dot"
	f, err := os.Create(fileName)

	if err != nil {
		t.Error("Can't create file: " + fileName)
	} else {
		_, err := f.WriteString(graph.GetDotRepresentation())
		if err != nil {
			t.Error("Can't write to  file: " + fileName)
		}
	}
	defer f.Close()

}
