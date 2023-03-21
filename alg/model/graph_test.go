package model

import (
	"testing"
)

func testGraphCreation(t *testing.T) {

	g := make(GraphRep)
	g[0] = []int{1, 2}
	g[1] = []int{0, 3}
	g[2] = []int{0, 3, 4}
	g[3] = []int{1, 2}
	g[4] = []int{2}
}
