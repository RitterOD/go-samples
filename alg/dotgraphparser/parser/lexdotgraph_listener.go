// Code generated from lexDotGraph.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // lexDotGraph

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// lexDotGraphListener is a complete listener for a parse tree produced by lexDotGraphParser.
type lexDotGraphListener interface {
	antlr.ParseTreeListener

	// EnterGraph is called when entering the graph production.
	EnterGraph(c *GraphContext)

	// EnterVertex_declaration is called when entering the vertex_declaration production.
	EnterVertex_declaration(c *Vertex_declarationContext)

	// EnterEdge_declaration is called when entering the edge_declaration production.
	EnterEdge_declaration(c *Edge_declarationContext)

	// EnterUndirect_edge_declaration is called when entering the undirect_edge_declaration production.
	EnterUndirect_edge_declaration(c *Undirect_edge_declarationContext)

	// EnterDirect_edge_declaration is called when entering the direct_edge_declaration production.
	EnterDirect_edge_declaration(c *Direct_edge_declarationContext)

	// EnterAttribute_declaration is called when entering the attribute_declaration production.
	EnterAttribute_declaration(c *Attribute_declarationContext)

	// ExitGraph is called when exiting the graph production.
	ExitGraph(c *GraphContext)

	// ExitVertex_declaration is called when exiting the vertex_declaration production.
	ExitVertex_declaration(c *Vertex_declarationContext)

	// ExitEdge_declaration is called when exiting the edge_declaration production.
	ExitEdge_declaration(c *Edge_declarationContext)

	// ExitUndirect_edge_declaration is called when exiting the undirect_edge_declaration production.
	ExitUndirect_edge_declaration(c *Undirect_edge_declarationContext)

	// ExitDirect_edge_declaration is called when exiting the direct_edge_declaration production.
	ExitDirect_edge_declaration(c *Direct_edge_declarationContext)

	// ExitAttribute_declaration is called when exiting the attribute_declaration production.
	ExitAttribute_declaration(c *Attribute_declarationContext)
}
