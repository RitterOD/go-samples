package graphgenerator

import (
	"fmt"
	"math"
	"math/rand"
	"sample/alg/model"
	"strconv"
	"time"
)

func ClusteredRandomGraph(minClusterSize, maxClusterSize int,
	minClusterCount, maxClusterCount int,
	clusterEdgeProbability, interClusterEdgeProbability float64,
	minWeight, maxWeight float64) *model.WeightedGraph {
	rv := model.NewWeightedGraph()
	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)
	clusterCount := minClusterCount + int(math.Round(randGenerator.Float64()*float64(maxClusterCount-minClusterCount)))
	currentClusterMinVertexIndex := 0
	currentClusterMaxVertexIndex := 0
	for i := 0; i < clusterCount; i++ {
		clusterSize := minClusterSize + int(math.Round(randGenerator.Float64()*float64(maxClusterSize-minClusterSize)))
		for j := 0; j < clusterSize; j++ {
			rv.AddVertex(currentClusterMinVertexIndex+j, strconv.Itoa(currentClusterMinVertexIndex+j))
		}
		currentClusterMaxVertexIndex += clusterSize
		for j := currentClusterMinVertexIndex; j < currentClusterMaxVertexIndex; j++ {
			for k := currentClusterMinVertexIndex; k < currentClusterMaxVertexIndex; k++ {
				if j != k {
					currentProb := rand.Float64()
					addEdge := currentProb > clusterEdgeProbability
					fmt.Println("currentProbability:" + fmt.Sprint(clusterSize))
					if addEdge {
						rv.AddEdge(j, k, minWeight+rand.Float64()*(maxWeight-minWeight))
						fmt.Println("edge added : " + fmt.Sprint(j) + " -> " + fmt.Sprint(k))
					}
				}
			}
		}
		currentClusterMinVertexIndex += clusterSize

	}
	return rv
}
