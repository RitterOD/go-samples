package main

import (
	"fmt"
	"sample/alg/model"
	"sample/alg/taskfield"
)

func main() {
	fmt.Println("Hello from alg")
	var filed = taskfield.GenerateIntField(10, 5, 1, 9)
	fmt.Println(filed)
	fmt.Println("Pretty print")
	taskfield.PrettyPrint(filed)
	var node1 = model.Create("1")
	var node2 = model.Create("2")
	var node3 = model.Create("3")
	model.AddLeft(node2, node1)
	model.AddRight(node2, node3)
	var dotRep = model.GetDotGraphStringRepresentation(node2)
	fmt.Printf("DOT REP\n")
	fmt.Println(dotRep)
	fmt.Printf("END DOT REP\n")
	var height = model.GetHeight(node2)
	fmt.Printf("Tree height: %d\n", height)

}
