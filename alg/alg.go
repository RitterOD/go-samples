package main

import (
	"fmt"
	"sample/alg/taskfield"
)

func main() {
	fmt.Println("Hello from alg")
	var filed = taskfield.GenerateIntField(10, 5, 1, 9)
	fmt.Println(filed)
	fmt.Println("Pretty print")
	taskfield.Pretty_print(filed)
}
