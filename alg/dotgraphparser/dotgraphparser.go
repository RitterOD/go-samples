package main

import (
	"fmt"
	"log"
	"os"
	"sample/alg/dotgraphparser/parser"
	"sample/alg/model"
)

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// based on https://github.com/antlr/antlr4/blob/master/doc/go-target.md
// based on https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/

type TreeShapeListener struct {
	*parser.BaselexDotGraphListener
}

type DotListener interface {
	antlr.ParseTreeListener
	EnterGraph(c *parser.GraphContext)
	ExitGraph(c *parser.GraphContext)
	EnterEdge_declaration(c *parser.Edge_declarationContext)
	ExitEdge_declaration(c *parser.Edge_declarationContext)
	EnterVertex_declaration(c *parser.Vertex_declarationContext)
	ExitVertex_declaration(c *parser.Vertex_declarationContext)
}

type dotListener struct {
	*parser.BaselexDotGraphListener
	text  string
	graph model.WeightedGraph
}

func (l *dotListener) EnterGraph(c *parser.GraphContext) {
	l.graph = *model.NewWeightedGraph()
}

func (l *dotListener) ExitGraph(c *parser.GraphContext) {
	l.graph.Name(c.NAME().GetText())
}

func (l *dotListener) EnterEdge_declaration(c *parser.Edge_declarationContext) {

}

func (l *dotListener) ExitEdge_declaration(c *parser.Edge_declarationContext) {
	fmt.Println("EDGE DECLARATION EXIT")
	if c.GetDirvar() != nil {
		fmt.Println("DIRECTED EDGE vertex1=" + c.GetDirvar().GetVname1().GetText() + " vertex2=" + c.GetDirvar().GetVname2().GetText())
	} else if c.GetUnvar() != nil {
		fmt.Println("UNDIRECTED EDGE vertex1=" + c.GetDirvar().GetVname1().GetText() + " vertex2=" + c.GetDirvar().GetVname2().GetText())
	}
	//fmt.Println(c.GetText())
}

func (l *dotListener) EnterVertex_declaration(c *parser.Vertex_declarationContext) {

}

func (l *dotListener) ExitVertex_declaration(c *parser.Vertex_declarationContext) {
	fmt.Println("VERTEX DECLARATION EXIT")
	//c.get
	fmt.Println(c.GetText())
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	fmt.Println("HELLO FROM DOT GRAPH PARSER PARSER")
	input, error := antlr.NewFileStream("./dotgraphparser/sample.dot")
	if error != nil {
		fmt.Println("ERROR OPEN FILE " + error.Error())
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path)
	}
	lexer := parser.NewlexDotGraphLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	var listener dotListener
	p := parser.NewlexDotGraphParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Graph())
	fmt.Println("GRAPH NAME = " + listener.graph.GetName())
	// LEXER USAGE DEMO
	//TO DO FIX
	//stream.GetAllText()
	//fmt.Println(lexer.GetSymbolicNames())
	//tokens := stream.GetAllTokens()
	//for _, t := range tokens {
	//	fmt.Println(t.GetText() + "type" + string(t.GetTokenType()))
	//}
}
