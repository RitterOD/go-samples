package model

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type binaryTreeNode struct {
	data   string
	parent *binaryTreeNode
	left   *binaryTreeNode
	right  *binaryTreeNode
}

func Create(data string) *binaryTreeNode {
	var rv binaryTreeNode
	rv.data = data
	return &rv
}

func AddRight(tree, node *binaryTreeNode) *binaryTreeNode {
	if tree != nil {
		tree.right = node
	}
	node.parent = tree
	return tree
}

func AddLeft(tree, node *binaryTreeNode) *binaryTreeNode {
	if tree != nil {
		tree.left = node
	}
	node.parent = tree
	return tree
}

func GetHeight(tree *binaryTreeNode) int64 {
	if tree == nil {
		return 0
	} else {
		leftHeight := GetHeight(tree.left)
		rightHeight := GetHeight(tree.right)
		if leftHeight > rightHeight {
			return 1 + leftHeight
		} else {
			return 1 + rightHeight
		}
	}
}

func GetDotGraphStringRepresentation(tree *binaryTreeNode) string {
	var nodes = make([]string, 0)
	var edges = make([]string, 0)
	getDotGraphStringRepresentationImpl(tree, &nodes, &edges)

	var rv = "digraph {" + strings.Join(nodes[:], ";\n") + ";\n"
	rv = rv + strings.Join(edges[:], ";\n") + ";\n}"
	return rv
}

func getDotGraphStringRepresentationImpl(tree *binaryTreeNode, nodes *[]string, edges *[]string) {
	if tree == nil {
		return
	} else {
		*nodes = append(*nodes, tree.data)
		if tree.left != nil {
			*edges = append(*edges, tree.data+" -> "+tree.left.data)
			getDotGraphStringRepresentationImpl(tree.left, nodes, edges)
		}
		if tree.right != nil {
			*edges = append(*edges, tree.data+" -> "+tree.right.data)
			getDotGraphStringRepresentationImpl(tree.right, nodes, edges)
		}
	}

}

func RandomTreeGenerator(height int64) *binaryTreeNode {
	rand.Seed(time.Now().UnixNano())
	return randomTreeGeneratorImpl(nil, height, 0, true)

}

func RandomStr(n int) string {
	//fmt.Printf("random string with %d", n)
	//rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:n]
}

func randomTreeGeneratorImpl(tree *binaryTreeNode, height, currentHeight int64, isLeft bool) *binaryTreeNode {
	var random = rand.Float64()
	if currentHeight < height {
		currentHeight += 1
	}
	if random/float64(currentHeight) > 0.10 {
		var node = Create("\"" + RandomStr(10) + "\"")
		if isLeft {
			AddLeft(tree, node)
		} else {
			AddRight(tree, node)
		}
		randomTreeGeneratorImpl(node, height, currentHeight, true)
		randomTreeGeneratorImpl(node, height, currentHeight, false)
		return node
	} else {
		return nil
	}

}
