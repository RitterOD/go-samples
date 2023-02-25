package taskfield

import (
	"fmt"
	"math/rand"
)

func GenerateIntField(height, width, lowBoundary, upBoundary int64) [][]int64 {
	var rv = make([][]int64, height)
	for i := int64(0); i < height; i++ {
		rv[i] = make([]int64, width)
		for j := int64(0); j < width; j++ {
			rv[i][j] = rand.Int63n(upBoundary-lowBoundary) + lowBoundary
		}
	}
	return rv
}

func Generate1dIntField(size, lowBoundary, upBoundary int64) []int64 {
	var rv = make([]int64, size)
	for i := int64(0); i < size; i++ {
		rv[i] = rand.Int63n(upBoundary-lowBoundary) + lowBoundary
	}
	return rv
}

func PrettyPrint(field [][]int64) {
	for _, i := range field {
		fmt.Println(i)
	}
}
