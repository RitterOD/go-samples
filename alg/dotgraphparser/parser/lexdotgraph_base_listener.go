// Code generated from lexDotGraph.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // lexDotGraph

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaselexDotGraphListener is a complete listener for a parse tree produced by lexDotGraphParser.
type BaselexDotGraphListener struct{}

var _ lexDotGraphListener = &BaselexDotGraphListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaselexDotGraphListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaselexDotGraphListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaselexDotGraphListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaselexDotGraphListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGraph is called when production graph is entered.
func (s *BaselexDotGraphListener) EnterGraph(ctx *GraphContext) {}

// ExitGraph is called when production graph is exited.
func (s *BaselexDotGraphListener) ExitGraph(ctx *GraphContext) {}

// EnterVertex_declaration is called when production vertex_declaration is entered.
func (s *BaselexDotGraphListener) EnterVertex_declaration(ctx *Vertex_declarationContext) {}

// ExitVertex_declaration is called when production vertex_declaration is exited.
func (s *BaselexDotGraphListener) ExitVertex_declaration(ctx *Vertex_declarationContext) {}

// EnterEdge_declaration is called when production edge_declaration is entered.
func (s *BaselexDotGraphListener) EnterEdge_declaration(ctx *Edge_declarationContext) {}

// ExitEdge_declaration is called when production edge_declaration is exited.
func (s *BaselexDotGraphListener) ExitEdge_declaration(ctx *Edge_declarationContext) {}

// EnterUndirect_edge_declaration is called when production undirect_edge_declaration is entered.
func (s *BaselexDotGraphListener) EnterUndirect_edge_declaration(ctx *Undirect_edge_declarationContext) {
}

// ExitUndirect_edge_declaration is called when production undirect_edge_declaration is exited.
func (s *BaselexDotGraphListener) ExitUndirect_edge_declaration(ctx *Undirect_edge_declarationContext) {
}

// EnterDirect_edge_declaration is called when production direct_edge_declaration is entered.
func (s *BaselexDotGraphListener) EnterDirect_edge_declaration(ctx *Direct_edge_declarationContext) {}

// ExitDirect_edge_declaration is called when production direct_edge_declaration is exited.
func (s *BaselexDotGraphListener) ExitDirect_edge_declaration(ctx *Direct_edge_declarationContext) {}

// EnterAttribute_declaration is called when production attribute_declaration is entered.
func (s *BaselexDotGraphListener) EnterAttribute_declaration(ctx *Attribute_declarationContext) {}

// ExitAttribute_declaration is called when production attribute_declaration is exited.
func (s *BaselexDotGraphListener) ExitAttribute_declaration(ctx *Attribute_declarationContext) {}
