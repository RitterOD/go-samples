// Code generated from lexDotGraph.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // lexDotGraph

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type lexDotGraphParser struct {
	*antlr.BaseParser
}

var lexdotgraphParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lexdotgraphParserInit() {
	staticData := &lexdotgraphParserStaticData
	staticData.literalNames = []string{
		"", "'digraph'", "'graph'", "", "'{'", "'}'", "'['", "']'", "'->'",
		"'=='", "'='", "';'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "NAME", "OPEN_BRACE", "CLOSE_BRACE", "OPEN_SQUARE_BRACKET",
		"CLOSE_SQUARE_BRACKET", "DIRECT_EDGE_OP", "UNDIRECT_EDGE_OP", "EQUALS_OP",
		"SEMICOLON", "WS",
	}
	staticData.ruleNames = []string{
		"graphType",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 5, 2, 0, 7, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 2,
		3, 0, 2, 1, 0, 0, 0, 2, 3, 7, 0, 0, 0, 3, 1, 1, 0, 0, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// lexDotGraphParserInit initializes any static state used to implement lexDotGraphParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewlexDotGraphParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LexDotGraphParserInit() {
	staticData := &lexdotgraphParserStaticData
	staticData.once.Do(lexdotgraphParserInit)
}

// NewlexDotGraphParser produces a new parser instance for the optional input antlr.TokenStream.
func NewlexDotGraphParser(input antlr.TokenStream) *lexDotGraphParser {
	LexDotGraphParserInit()
	this := new(lexDotGraphParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &lexdotgraphParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "lexDotGraph.g4"

	return this
}

// lexDotGraphParser tokens.
const (
	lexDotGraphParserEOF                  = antlr.TokenEOF
	lexDotGraphParserT__0                 = 1
	lexDotGraphParserT__1                 = 2
	lexDotGraphParserNAME                 = 3
	lexDotGraphParserOPEN_BRACE           = 4
	lexDotGraphParserCLOSE_BRACE          = 5
	lexDotGraphParserOPEN_SQUARE_BRACKET  = 6
	lexDotGraphParserCLOSE_SQUARE_BRACKET = 7
	lexDotGraphParserDIRECT_EDGE_OP       = 8
	lexDotGraphParserUNDIRECT_EDGE_OP     = 9
	lexDotGraphParserEQUALS_OP            = 10
	lexDotGraphParserSEMICOLON            = 11
	lexDotGraphParserWS                   = 12
)

// lexDotGraphParserRULE_graphType is the lexDotGraphParser rule.
const lexDotGraphParserRULE_graphType = 0

// IGraphTypeContext is an interface to support dynamic dispatch.
type IGraphTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsGraphTypeContext differentiates from other interfaces.
	IsGraphTypeContext()
}

type GraphTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphTypeContext() *GraphTypeContext {
	var p = new(GraphTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lexDotGraphParserRULE_graphType
	return p
}

func (*GraphTypeContext) IsGraphTypeContext() {}

func NewGraphTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphTypeContext {
	var p = new(GraphTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lexDotGraphParserRULE_graphType

	return p
}

func (s *GraphTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *GraphTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.EnterGraphType(s)
	}
}

func (s *GraphTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.ExitGraphType(s)
	}
}

func (p *lexDotGraphParser) GraphType() (localctx IGraphTypeContext) {
	this := p
	_ = this

	localctx = NewGraphTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, lexDotGraphParserRULE_graphType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(2)
		_la = p.GetTokenStream().LA(1)

		if !(_la == lexDotGraphParserT__0 || _la == lexDotGraphParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
