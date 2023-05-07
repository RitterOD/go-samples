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
		"", "", "", "'{'", "'}'", "'['", "']'", "'->'", "'=='", "'='", "';'",
	}
	staticData.symbolicNames = []string{
		"", "GRAPH_TYPE", "NAME", "OPEN_BRACE", "CLOSE_BRACE", "OPEN_SQUARE_BRACKET",
		"CLOSE_SQUARE_BRACKET", "DIRECT_EDGE_OP", "UNDIRECT_EDGE_OP", "EQUALS_OP",
		"SEMICOLON", "WS",
	}
	staticData.ruleNames = []string{
		"graph", "vertex_declaration", "edge_declaration", "undirect_edge_declaration",
		"direct_edge_declaration", "attribute_declaration",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 11, 55, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 18, 8, 0, 10, 0, 12,
		0, 21, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 3, 2, 34, 8, 2, 1, 2, 3, 2, 37, 8, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 0, 0, 6, 0, 2, 4, 6, 8, 10, 0, 0, 53, 0, 12, 1, 0, 0, 0, 2, 25,
		1, 0, 0, 0, 4, 33, 1, 0, 0, 0, 6, 40, 1, 0, 0, 0, 8, 44, 1, 0, 0, 0, 10,
		48, 1, 0, 0, 0, 12, 13, 5, 1, 0, 0, 13, 14, 5, 2, 0, 0, 14, 19, 5, 3, 0,
		0, 15, 18, 3, 2, 1, 0, 16, 18, 3, 4, 2, 0, 17, 15, 1, 0, 0, 0, 17, 16,
		1, 0, 0, 0, 18, 21, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0,
		20, 22, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 22, 23, 5, 4, 0, 0, 23, 24, 5,
		0, 0, 1, 24, 1, 1, 0, 0, 0, 25, 27, 5, 2, 0, 0, 26, 28, 3, 10, 5, 0, 27,
		26, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 30, 5, 10,
		0, 0, 30, 3, 1, 0, 0, 0, 31, 34, 3, 6, 3, 0, 32, 34, 3, 8, 4, 0, 33, 31,
		1, 0, 0, 0, 33, 32, 1, 0, 0, 0, 34, 36, 1, 0, 0, 0, 35, 37, 3, 10, 5, 0,
		36, 35, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 39, 5,
		10, 0, 0, 39, 5, 1, 0, 0, 0, 40, 41, 5, 2, 0, 0, 41, 42, 5, 8, 0, 0, 42,
		43, 5, 2, 0, 0, 43, 7, 1, 0, 0, 0, 44, 45, 5, 2, 0, 0, 45, 46, 5, 7, 0,
		0, 46, 47, 5, 2, 0, 0, 47, 9, 1, 0, 0, 0, 48, 49, 5, 5, 0, 0, 49, 50, 5,
		2, 0, 0, 50, 51, 5, 9, 0, 0, 51, 52, 5, 2, 0, 0, 52, 53, 5, 6, 0, 0, 53,
		11, 1, 0, 0, 0, 5, 17, 19, 27, 33, 36,
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
	lexDotGraphParserGRAPH_TYPE           = 1
	lexDotGraphParserNAME                 = 2
	lexDotGraphParserOPEN_BRACE           = 3
	lexDotGraphParserCLOSE_BRACE          = 4
	lexDotGraphParserOPEN_SQUARE_BRACKET  = 5
	lexDotGraphParserCLOSE_SQUARE_BRACKET = 6
	lexDotGraphParserDIRECT_EDGE_OP       = 7
	lexDotGraphParserUNDIRECT_EDGE_OP     = 8
	lexDotGraphParserEQUALS_OP            = 9
	lexDotGraphParserSEMICOLON            = 10
	lexDotGraphParserWS                   = 11
)

// lexDotGraphParser rules.
const (
	lexDotGraphParserRULE_graph                     = 0
	lexDotGraphParserRULE_vertex_declaration        = 1
	lexDotGraphParserRULE_edge_declaration          = 2
	lexDotGraphParserRULE_undirect_edge_declaration = 3
	lexDotGraphParserRULE_direct_edge_declaration   = 4
	lexDotGraphParserRULE_attribute_declaration     = 5
)

// IGraphContext is an interface to support dynamic dispatch.
type IGraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GRAPH_TYPE() antlr.TerminalNode
	NAME() antlr.TerminalNode
	OPEN_BRACE() antlr.TerminalNode
	CLOSE_BRACE() antlr.TerminalNode
	EOF() antlr.TerminalNode
	AllVertex_declaration() []IVertex_declarationContext
	Vertex_declaration(i int) IVertex_declarationContext
	AllEdge_declaration() []IEdge_declarationContext
	Edge_declaration(i int) IEdge_declarationContext

	// IsGraphContext differentiates from other interfaces.
	IsGraphContext()
}

type GraphContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphContext() *GraphContext {
	var p = new(GraphContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lexDotGraphParserRULE_graph
	return p
}

func (*GraphContext) IsGraphContext() {}

func NewGraphContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphContext {
	var p = new(GraphContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lexDotGraphParserRULE_graph

	return p
}

func (s *GraphContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphContext) GRAPH_TYPE() antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserGRAPH_TYPE, 0)
}

func (s *GraphContext) NAME() antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserNAME, 0)
}

func (s *GraphContext) OPEN_BRACE() antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserOPEN_BRACE, 0)
}

func (s *GraphContext) CLOSE_BRACE() antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserCLOSE_BRACE, 0)
}

func (s *GraphContext) EOF() antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserEOF, 0)
}

func (s *GraphContext) AllVertex_declaration() []IVertex_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVertex_declarationContext); ok {
			len++
		}
	}

	tst := make([]IVertex_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVertex_declarationContext); ok {
			tst[i] = t.(IVertex_declarationContext)
			i++
		}
	}

	return tst
}

func (s *GraphContext) Vertex_declaration(i int) IVertex_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVertex_declarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVertex_declarationContext)
}

func (s *GraphContext) AllEdge_declaration() []IEdge_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEdge_declarationContext); ok {
			len++
		}
	}

	tst := make([]IEdge_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEdge_declarationContext); ok {
			tst[i] = t.(IEdge_declarationContext)
			i++
		}
	}

	return tst
}

func (s *GraphContext) Edge_declaration(i int) IEdge_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEdge_declarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEdge_declarationContext)
}

func (s *GraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.EnterGraph(s)
	}
}

func (s *GraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.ExitGraph(s)
	}
}

func (p *lexDotGraphParser) Graph() (localctx IGraphContext) {
	this := p
	_ = this

	localctx = NewGraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, lexDotGraphParserRULE_graph)
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
		p.SetState(12)
		p.Match(lexDotGraphParserGRAPH_TYPE)
	}
	{
		p.SetState(13)
		p.Match(lexDotGraphParserNAME)
	}
	{
		p.SetState(14)
		p.Match(lexDotGraphParserOPEN_BRACE)
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == lexDotGraphParserNAME {
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(15)
				p.Vertex_declaration()
			}

		case 2:
			{
				p.SetState(16)
				p.Edge_declaration()
			}

		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(22)
		p.Match(lexDotGraphParserCLOSE_BRACE)
	}
	{
		p.SetState(23)
		p.Match(lexDotGraphParserEOF)
	}

	return localctx
}

// IVertex_declarationContext is an interface to support dynamic dispatch.
type IVertex_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVname returns the vname token.
	GetVname() antlr.Token

	// SetVname sets the vname token.
	SetVname(antlr.Token)

	// Getter signatures
	SEMICOLON() antlr.TerminalNode
	NAME() antlr.TerminalNode
	Attribute_declaration() IAttribute_declarationContext

	// IsVertex_declarationContext differentiates from other interfaces.
	IsVertex_declarationContext()
}

type Vertex_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	vname  antlr.Token
}

func NewEmptyVertex_declarationContext() *Vertex_declarationContext {
	var p = new(Vertex_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lexDotGraphParserRULE_vertex_declaration
	return p
}

func (*Vertex_declarationContext) IsVertex_declarationContext() {}

func NewVertex_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vertex_declarationContext {
	var p = new(Vertex_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lexDotGraphParserRULE_vertex_declaration

	return p
}

func (s *Vertex_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Vertex_declarationContext) GetVname() antlr.Token { return s.vname }

func (s *Vertex_declarationContext) SetVname(v antlr.Token) { s.vname = v }

func (s *Vertex_declarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserSEMICOLON, 0)
}

func (s *Vertex_declarationContext) NAME() antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserNAME, 0)
}

func (s *Vertex_declarationContext) Attribute_declaration() IAttribute_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttribute_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttribute_declarationContext)
}

func (s *Vertex_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vertex_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vertex_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.EnterVertex_declaration(s)
	}
}

func (s *Vertex_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.ExitVertex_declaration(s)
	}
}

func (p *lexDotGraphParser) Vertex_declaration() (localctx IVertex_declarationContext) {
	this := p
	_ = this

	localctx = NewVertex_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, lexDotGraphParserRULE_vertex_declaration)
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
		p.SetState(25)

		var _m = p.Match(lexDotGraphParserNAME)

		localctx.(*Vertex_declarationContext).vname = _m
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == lexDotGraphParserOPEN_SQUARE_BRACKET {
		{
			p.SetState(26)
			p.Attribute_declaration()
		}

	}
	{
		p.SetState(29)
		p.Match(lexDotGraphParserSEMICOLON)
	}

	return localctx
}

// IEdge_declarationContext is an interface to support dynamic dispatch.
type IEdge_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetUnvar returns the unvar rule contexts.
	GetUnvar() IUndirect_edge_declarationContext

	// GetDirvar returns the dirvar rule contexts.
	GetDirvar() IDirect_edge_declarationContext

	// SetUnvar sets the unvar rule contexts.
	SetUnvar(IUndirect_edge_declarationContext)

	// SetDirvar sets the dirvar rule contexts.
	SetDirvar(IDirect_edge_declarationContext)

	// Getter signatures
	SEMICOLON() antlr.TerminalNode
	Undirect_edge_declaration() IUndirect_edge_declarationContext
	Direct_edge_declaration() IDirect_edge_declarationContext
	Attribute_declaration() IAttribute_declarationContext

	// IsEdge_declarationContext differentiates from other interfaces.
	IsEdge_declarationContext()
}

type Edge_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	unvar  IUndirect_edge_declarationContext
	dirvar IDirect_edge_declarationContext
}

func NewEmptyEdge_declarationContext() *Edge_declarationContext {
	var p = new(Edge_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lexDotGraphParserRULE_edge_declaration
	return p
}

func (*Edge_declarationContext) IsEdge_declarationContext() {}

func NewEdge_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Edge_declarationContext {
	var p = new(Edge_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lexDotGraphParserRULE_edge_declaration

	return p
}

func (s *Edge_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Edge_declarationContext) GetUnvar() IUndirect_edge_declarationContext { return s.unvar }

func (s *Edge_declarationContext) GetDirvar() IDirect_edge_declarationContext { return s.dirvar }

func (s *Edge_declarationContext) SetUnvar(v IUndirect_edge_declarationContext) { s.unvar = v }

func (s *Edge_declarationContext) SetDirvar(v IDirect_edge_declarationContext) { s.dirvar = v }

func (s *Edge_declarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserSEMICOLON, 0)
}

func (s *Edge_declarationContext) Undirect_edge_declaration() IUndirect_edge_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUndirect_edge_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUndirect_edge_declarationContext)
}

func (s *Edge_declarationContext) Direct_edge_declaration() IDirect_edge_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirect_edge_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirect_edge_declarationContext)
}

func (s *Edge_declarationContext) Attribute_declaration() IAttribute_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttribute_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttribute_declarationContext)
}

func (s *Edge_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Edge_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Edge_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.EnterEdge_declaration(s)
	}
}

func (s *Edge_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.ExitEdge_declaration(s)
	}
}

func (p *lexDotGraphParser) Edge_declaration() (localctx IEdge_declarationContext) {
	this := p
	_ = this

	localctx = NewEdge_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, lexDotGraphParserRULE_edge_declaration)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(31)

			var _x = p.Undirect_edge_declaration()

			localctx.(*Edge_declarationContext).unvar = _x
		}

	case 2:
		{
			p.SetState(32)

			var _x = p.Direct_edge_declaration()

			localctx.(*Edge_declarationContext).dirvar = _x
		}

	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == lexDotGraphParserOPEN_SQUARE_BRACKET {
		{
			p.SetState(35)
			p.Attribute_declaration()
		}

	}
	{
		p.SetState(38)
		p.Match(lexDotGraphParserSEMICOLON)
	}

	return localctx
}

// IUndirect_edge_declarationContext is an interface to support dynamic dispatch.
type IUndirect_edge_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVname1 returns the vname1 token.
	GetVname1() antlr.Token

	// GetVname2 returns the vname2 token.
	GetVname2() antlr.Token

	// SetVname1 sets the vname1 token.
	SetVname1(antlr.Token)

	// SetVname2 sets the vname2 token.
	SetVname2(antlr.Token)

	// Getter signatures
	UNDIRECT_EDGE_OP() antlr.TerminalNode
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode

	// IsUndirect_edge_declarationContext differentiates from other interfaces.
	IsUndirect_edge_declarationContext()
}

type Undirect_edge_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	vname1 antlr.Token
	vname2 antlr.Token
}

func NewEmptyUndirect_edge_declarationContext() *Undirect_edge_declarationContext {
	var p = new(Undirect_edge_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lexDotGraphParserRULE_undirect_edge_declaration
	return p
}

func (*Undirect_edge_declarationContext) IsUndirect_edge_declarationContext() {}

func NewUndirect_edge_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Undirect_edge_declarationContext {
	var p = new(Undirect_edge_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lexDotGraphParserRULE_undirect_edge_declaration

	return p
}

func (s *Undirect_edge_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Undirect_edge_declarationContext) GetVname1() antlr.Token { return s.vname1 }

func (s *Undirect_edge_declarationContext) GetVname2() antlr.Token { return s.vname2 }

func (s *Undirect_edge_declarationContext) SetVname1(v antlr.Token) { s.vname1 = v }

func (s *Undirect_edge_declarationContext) SetVname2(v antlr.Token) { s.vname2 = v }

func (s *Undirect_edge_declarationContext) UNDIRECT_EDGE_OP() antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserUNDIRECT_EDGE_OP, 0)
}

func (s *Undirect_edge_declarationContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(lexDotGraphParserNAME)
}

func (s *Undirect_edge_declarationContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserNAME, i)
}

func (s *Undirect_edge_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Undirect_edge_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Undirect_edge_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.EnterUndirect_edge_declaration(s)
	}
}

func (s *Undirect_edge_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.ExitUndirect_edge_declaration(s)
	}
}

func (p *lexDotGraphParser) Undirect_edge_declaration() (localctx IUndirect_edge_declarationContext) {
	this := p
	_ = this

	localctx = NewUndirect_edge_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, lexDotGraphParserRULE_undirect_edge_declaration)

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
		p.SetState(40)

		var _m = p.Match(lexDotGraphParserNAME)

		localctx.(*Undirect_edge_declarationContext).vname1 = _m
	}
	{
		p.SetState(41)
		p.Match(lexDotGraphParserUNDIRECT_EDGE_OP)
	}
	{
		p.SetState(42)

		var _m = p.Match(lexDotGraphParserNAME)

		localctx.(*Undirect_edge_declarationContext).vname2 = _m
	}

	return localctx
}

// IDirect_edge_declarationContext is an interface to support dynamic dispatch.
type IDirect_edge_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVname1 returns the vname1 token.
	GetVname1() antlr.Token

	// GetVname2 returns the vname2 token.
	GetVname2() antlr.Token

	// SetVname1 sets the vname1 token.
	SetVname1(antlr.Token)

	// SetVname2 sets the vname2 token.
	SetVname2(antlr.Token)

	// Getter signatures
	DIRECT_EDGE_OP() antlr.TerminalNode
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode

	// IsDirect_edge_declarationContext differentiates from other interfaces.
	IsDirect_edge_declarationContext()
}

type Direct_edge_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	vname1 antlr.Token
	vname2 antlr.Token
}

func NewEmptyDirect_edge_declarationContext() *Direct_edge_declarationContext {
	var p = new(Direct_edge_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lexDotGraphParserRULE_direct_edge_declaration
	return p
}

func (*Direct_edge_declarationContext) IsDirect_edge_declarationContext() {}

func NewDirect_edge_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Direct_edge_declarationContext {
	var p = new(Direct_edge_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lexDotGraphParserRULE_direct_edge_declaration

	return p
}

func (s *Direct_edge_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Direct_edge_declarationContext) GetVname1() antlr.Token { return s.vname1 }

func (s *Direct_edge_declarationContext) GetVname2() antlr.Token { return s.vname2 }

func (s *Direct_edge_declarationContext) SetVname1(v antlr.Token) { s.vname1 = v }

func (s *Direct_edge_declarationContext) SetVname2(v antlr.Token) { s.vname2 = v }

func (s *Direct_edge_declarationContext) DIRECT_EDGE_OP() antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserDIRECT_EDGE_OP, 0)
}

func (s *Direct_edge_declarationContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(lexDotGraphParserNAME)
}

func (s *Direct_edge_declarationContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserNAME, i)
}

func (s *Direct_edge_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Direct_edge_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Direct_edge_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.EnterDirect_edge_declaration(s)
	}
}

func (s *Direct_edge_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.ExitDirect_edge_declaration(s)
	}
}

func (p *lexDotGraphParser) Direct_edge_declaration() (localctx IDirect_edge_declarationContext) {
	this := p
	_ = this

	localctx = NewDirect_edge_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, lexDotGraphParserRULE_direct_edge_declaration)

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
		p.SetState(44)

		var _m = p.Match(lexDotGraphParserNAME)

		localctx.(*Direct_edge_declarationContext).vname1 = _m
	}
	{
		p.SetState(45)
		p.Match(lexDotGraphParserDIRECT_EDGE_OP)
	}
	{
		p.SetState(46)

		var _m = p.Match(lexDotGraphParserNAME)

		localctx.(*Direct_edge_declarationContext).vname2 = _m
	}

	return localctx
}

// IAttribute_declarationContext is an interface to support dynamic dispatch.
type IAttribute_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAttName returns the attName token.
	GetAttName() antlr.Token

	// GetAttValue returns the attValue token.
	GetAttValue() antlr.Token

	// SetAttName sets the attName token.
	SetAttName(antlr.Token)

	// SetAttValue sets the attValue token.
	SetAttValue(antlr.Token)

	// Getter signatures
	OPEN_SQUARE_BRACKET() antlr.TerminalNode
	EQUALS_OP() antlr.TerminalNode
	CLOSE_SQUARE_BRACKET() antlr.TerminalNode
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode

	// IsAttribute_declarationContext differentiates from other interfaces.
	IsAttribute_declarationContext()
}

type Attribute_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	attName  antlr.Token
	attValue antlr.Token
}

func NewEmptyAttribute_declarationContext() *Attribute_declarationContext {
	var p = new(Attribute_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lexDotGraphParserRULE_attribute_declaration
	return p
}

func (*Attribute_declarationContext) IsAttribute_declarationContext() {}

func NewAttribute_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Attribute_declarationContext {
	var p = new(Attribute_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lexDotGraphParserRULE_attribute_declaration

	return p
}

func (s *Attribute_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Attribute_declarationContext) GetAttName() antlr.Token { return s.attName }

func (s *Attribute_declarationContext) GetAttValue() antlr.Token { return s.attValue }

func (s *Attribute_declarationContext) SetAttName(v antlr.Token) { s.attName = v }

func (s *Attribute_declarationContext) SetAttValue(v antlr.Token) { s.attValue = v }

func (s *Attribute_declarationContext) OPEN_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserOPEN_SQUARE_BRACKET, 0)
}

func (s *Attribute_declarationContext) EQUALS_OP() antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserEQUALS_OP, 0)
}

func (s *Attribute_declarationContext) CLOSE_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserCLOSE_SQUARE_BRACKET, 0)
}

func (s *Attribute_declarationContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(lexDotGraphParserNAME)
}

func (s *Attribute_declarationContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(lexDotGraphParserNAME, i)
}

func (s *Attribute_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Attribute_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Attribute_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.EnterAttribute_declaration(s)
	}
}

func (s *Attribute_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lexDotGraphListener); ok {
		listenerT.ExitAttribute_declaration(s)
	}
}

func (p *lexDotGraphParser) Attribute_declaration() (localctx IAttribute_declarationContext) {
	this := p
	_ = this

	localctx = NewAttribute_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, lexDotGraphParserRULE_attribute_declaration)

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
		p.SetState(48)
		p.Match(lexDotGraphParserOPEN_SQUARE_BRACKET)
	}
	{
		p.SetState(49)

		var _m = p.Match(lexDotGraphParserNAME)

		localctx.(*Attribute_declarationContext).attName = _m
	}
	{
		p.SetState(50)
		p.Match(lexDotGraphParserEQUALS_OP)
	}
	{
		p.SetState(51)

		var _m = p.Match(lexDotGraphParserNAME)

		localctx.(*Attribute_declarationContext).attValue = _m
	}
	{
		p.SetState(52)
		p.Match(lexDotGraphParserCLOSE_SQUARE_BRACKET)
	}

	return localctx
}
