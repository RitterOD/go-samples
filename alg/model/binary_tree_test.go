package model

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestSerializationDeserialization(t *testing.T) {
	var node1 = Create("1")
	var node2 = Create("2")
	var node3 = Create("3")
	AddLeft(node2, node1)
	AddRight(node2, node3)
	data, err := json.MarshalIndent(node2, "", "")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("Serialized tree:\n%s\n", data)
	var node BinaryTreeNode
	{
	}
	json.Unmarshal(data, &node)

	if node.Data != "2" {
		t.Error("FAIL TO RIGHT UNMARSHAL")
	}

}
