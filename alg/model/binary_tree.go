package model

import "strings"

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
	tree.right = node
	node.parent = tree
	return tree
}

func AddLeft(tree, node *binaryTreeNode) *binaryTreeNode {
	tree.left = node
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
