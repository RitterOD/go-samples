package graphgenerator

import (
	"math"
	"math/rand"
	"sample/alg/model"
	"strconv"
	"time"
)

func ClusteredRandomGraph(minClusterSize, maxClusterSize int,
	minClusterCount, maxClusterCount int,
	clusterEdgeProbability, interClusterEdgeProbability float64,
	minWeight, maxWeight float64, interClusterMinWeight, interClusterMaxWeight float64) *model.WeightedGraph {
	rv := model.NewWeightedGraph()
	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)
	clusterCount := minClusterCount + int(math.Round(randGenerator.Float64()*float64(maxClusterCount-minClusterCount)))
	currentClusterMinVertexIndex := 0
	currentClusterMaxVertexIndex := 0
	clusterIdToVertexRange := make(map[int][]int)
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
					addEdge := currentProb < clusterEdgeProbability
					//fmt.Println("currentProbability:" + fmt.Sprint(clusterSize))
					if addEdge {
						rv.AddEdge(j, k, minWeight+rand.Float64()*(maxWeight-minWeight))
						//fmt.Println("edge added : " + fmt.Sprint(j) + " -> " + fmt.Sprint(k))
					}
				}
			}
		}
		clusterIdToVertexRange[i] = []int{currentClusterMinVertexIndex, currentClusterMaxVertexIndex}
		currentClusterMinVertexIndex += clusterSize

	}
	for _, r := range clusterIdToVertexRange {
		for _, rTo := range clusterIdToVertexRange {
			if r[0] != rTo[0] {
				currentProb := rand.Float64()
				if currentProb < interClusterEdgeProbability {
					outVertexInd := r[0] + int(math.Round(float64(r[1]-r[0])*rand.Float64()))
					inVertexInd := rTo[0] + int(math.Round(float64(rTo[1]-rTo[0])*rand.Float64()))
					rv.AddEdge(outVertexInd, inVertexInd, interClusterMinWeight+
						rand.Float64()*(interClusterMaxWeight-interClusterMinWeight))
				}
			}
		}
	}
	return rv
}
