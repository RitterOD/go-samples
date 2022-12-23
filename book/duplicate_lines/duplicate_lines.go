package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, error := os.ReadFile(filename)
		if error != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", error)
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		for line, n := range counts {
			fmt.Printf("%s\t%d\n", line, n)
		}

	}
}
