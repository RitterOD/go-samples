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

// EnterGraphType is called when production graphType is entered.
func (s *BaselexDotGraphListener) EnterGraphType(ctx *GraphTypeContext) {}

// ExitGraphType is called when production graphType is exited.
func (s *BaselexDotGraphListener) ExitGraphType(ctx *GraphTypeContext) {}
