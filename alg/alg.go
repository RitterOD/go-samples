package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	data, err := json.MarshalIndent(node2, "", "")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("Serialized tree:\n%s\n", data)
	var node model.BinaryTreeNode
	{
	}
	json.Unmarshal(data, &node)
	fmt.Printf("Data of unmarshalled json: " + node.Data)
	var height = model.GetHeight(node2)
	fmt.Printf("Tree height: %d\n", height)

	var tree = model.RandomTreeGenerator(10, 0.05)
	if tree != nil {
		dotRep = model.GetDotGraphStringRepresentation(tree)
		//fmt.Printf("DOT REP RANDOM TREE\n")
		//fmt.Println(dotRep)
		//fmt.Printf("END DOT REP RANDOM TREE\n")
		f, err := os.Create("btree.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer f.Close()
		_, err = f.WriteString(dotRep)
		if err != nil {
			log.Fatal(err)
			return
		}

	}

}
