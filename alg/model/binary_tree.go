package model

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type BinaryTreeNode struct {
	Data   string
	parent *BinaryTreeNode
	Left   *BinaryTreeNode
	Right  *BinaryTreeNode
}

func Create(data string) *BinaryTreeNode {
	var rv BinaryTreeNode
	rv.Data = data
	return &rv
}

func AddRight(tree, node *BinaryTreeNode) *BinaryTreeNode {
	if tree != nil {
		tree.Right = node
	}
	node.parent = tree
	return tree
}

func AddLeft(tree, node *BinaryTreeNode) *BinaryTreeNode {
	if tree != nil {
		tree.Left = node
	}
	node.parent = tree
	return tree
}

func GetHeight(tree *BinaryTreeNode) int64 {
	if tree == nil {
		return 0
	} else {
		leftHeight := GetHeight(tree.Left)
		rightHeight := GetHeight(tree.Right)
		if leftHeight > rightHeight {
			return 1 + leftHeight
		} else {
			return 1 + rightHeight
		}
	}
}

func GetDotGraphStringRepresentation(tree *BinaryTreeNode) string {
	var nodes = make([]string, 0)
	var edges = make([]string, 0)
	getDotGraphStringRepresentationImpl(tree, &nodes, &edges)

	var rv = "digraph {" + strings.Join(nodes[:], ";\n") + ";\n"
	rv = rv + strings.Join(edges[:], ";\n") + ";\n}"
	return rv
}

func getDotGraphStringRepresentationImpl(tree *BinaryTreeNode, nodes *[]string, edges *[]string) {
	if tree == nil {
		return
	} else {
		*nodes = append(*nodes, tree.Data)
		if tree.Left != nil {
			*edges = append(*edges, tree.Data+" -> "+tree.Left.Data)
			getDotGraphStringRepresentationImpl(tree.Left, nodes, edges)
		}
		if tree.Right != nil {
			*edges = append(*edges, tree.Data+" -> "+tree.Right.Data)
			getDotGraphStringRepresentationImpl(tree.Right, nodes, edges)
		}
	}

}

func RandomTreeGenerator(height int64, threshold float64) *BinaryTreeNode {
	rand.Seed(time.Now().UnixNano())
	return randomTreeGeneratorImpl(nil, height, 0, threshold, true)

}

func RandomStr(n int) string {
	//fmt.Printf("random string with %d", n)
	//rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:n]
}

func randomTreeGeneratorImpl(tree *BinaryTreeNode, height, currentHeight int64, threshold float64, isLeft bool) *BinaryTreeNode {
	var random = rand.Float64()
	if currentHeight < height {
		currentHeight += 1
	}
	if random/float64(currentHeight) > threshold {
		var node = Create("\"" + RandomStr(10) + "\"")
		if isLeft {
			AddLeft(tree, node)
		} else {
			AddRight(tree, node)
		}
		randomTreeGeneratorImpl(node, height, currentHeight, threshold, true)
		randomTreeGeneratorImpl(node, height, currentHeight, threshold, false)
		return node
	} else {
		return nil
	}

}
