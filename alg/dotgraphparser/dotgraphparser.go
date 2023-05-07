package main

import (
	"fmt"
	"log"
	"os"
	"sample/alg/dotgraphparser/parser"
	"sample/alg/model"
	"strconv"
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
	text               string
	graph              model.WeightedGraph
	currentVertexIndex int
	currentWeight      float64
}

func (l *dotListener) EnterGraph(c *parser.GraphContext) {
	l.graph = *model.NewWeightedGraph()
}

func (l *dotListener) ExitGraph(c *parser.GraphContext) {
	l.graph.Name(c.NAME().GetText())
}

func (l *dotListener) EnterEdge_declaration(c *parser.Edge_declarationContext) {

}

func (l *dotListener) ExitAttribute_declaration(c *parser.Attribute_declarationContext) {
	txtvalue := c.GetAttValue().GetText()
	weight, err := strconv.ParseFloat(txtvalue, 64)
	if err != nil {
		l.currentWeight = 0
		fmt.Println("cant convert to float" + txtvalue)
	} else {
		l.currentWeight = weight
	}

}

// TODO ADD ERROR HANDLING
func (l *dotListener) ExitEdge_declaration(c *parser.Edge_declarationContext) {
	if c.GetDirvar() != nil {
		//fmt.Println("DIRECTED EDGE vertex1=" + c.GetDirvar().GetVname1().GetText() + " vertex2=" + c.GetDirvar().GetVname2().GetText())
		l.graph.AddEdgeByVertexName(c.GetDirvar().GetVname1().GetText(), c.GetDirvar().GetVname2().GetText(), l.currentWeight)
	} else if c.GetUnvar() != nil {
		panic("UNSUPPORTED TYPE OF EDGE")
	}
}

func (l *dotListener) EnterVertex_declaration(c *parser.Vertex_declarationContext) {

}

func (l *dotListener) ExitVertex_declaration(c *parser.Vertex_declarationContext) {
	l.graph.AddVertex(l.currentVertexIndex, c.NAME().GetText())
	l.currentVertexIndex++
	//fmt.Println(c.GetText())
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
	fmt.Println("END OF PARSING")
	// LEXER USAGE DEMO
	//TO DO FIX
	//stream.GetAllText()
	//fmt.Println(lexer.GetSymbolicNames())
	//tokens := stream.GetAllTokens()
	//for _, t := range tokens {
	//	fmt.Println(t.GetText() + "type" + string(t.GetTokenType()))
	//}
}
