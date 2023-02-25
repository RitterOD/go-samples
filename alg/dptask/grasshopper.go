package main

import (
	"fmt"
	"sample/alg/taskfield"
)

func calculateNumberOfWays(length int64) int64 {
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 1
	}
	if length == 2 {
		return 2
	}
	var step1 = int64(1)
	var step2 = int64(2)
	var tmp = int64(3)
	for curStep := int64(2); curStep < length; curStep++ {
		tmp = step1 + step2
		step1 = step2
		step2 = tmp
	}
	return tmp
}

func main() {
	var field = taskfield.Generate1dIntField(5, -10, 10)
	fmt.Println("task field")
	fmt.Println(field)
	var rv = calculateNumberOfWays(1)
	fmt.Printf("length: %d number of ways: %d\n", 1, rv)
	rv = calculateNumberOfWays(2)
	fmt.Printf("length: %d number of ways: %d\n", 2, rv)
	for i := 3; i < 100; i++ {
		rv = calculateNumberOfWays(int64(i))
		fmt.Printf("length: %d number of ways: %d\n", i, rv)
	}
}
