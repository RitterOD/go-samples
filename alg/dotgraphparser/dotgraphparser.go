package main

import (
	"fmt"
	"log"
	"os"
	"sample/alg/dotgraphparser/parser"
)

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// based on https://github.com/antlr/antlr4/blob/master/doc/go-target.md

type TreeShapeListener struct {
	*parser.BaselexDotGraphListener
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
	stream.GetAllText()
	fmt.Println(lexer.GetSymbolicNames())
	tokens := stream.GetAllTokens()
	for _, t := range tokens {
		fmt.Println(t.GetText() + "type" + string(t.GetTokenType()))
	}
}
