// Code generated from lexDotGraph.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type lexDotGraphLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexdotgraphlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lexdotgraphlexerLexerInit() {
	staticData := &lexdotgraphlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "", "'{'", "'}'", "'['", "']'", "'->'", "'=='", "'='", "';'",
	}
	staticData.symbolicNames = []string{
		"", "GRAPH_TYPE", "NAME", "OPEN_BRACE", "CLOSE_BRACE", "OPEN_SQUARE_BRACKET",
		"CLOSE_SQUARE_BRACKET", "DIRECT_EDGE_OP", "UNDIRECT_EDGE_OP", "EQUALS_OP",
		"SEMICOLON", "WS",
	}
	staticData.ruleNames = []string{
		"GRAPH_TYPE", "NAME", "OPEN_BRACE", "CLOSE_BRACE", "OPEN_SQUARE_BRACKET",
		"CLOSE_SQUARE_BRACKET", "DIRECT_EDGE_OP", "UNDIRECT_EDGE_OP", "EQUALS_OP",
		"SEMICOLON", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 11, 67, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 3, 0, 36, 8, 0, 1, 1, 4, 1, 39, 8, 1, 11, 1, 12, 1, 40, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 4, 10, 62, 8, 10, 11, 10, 12, 10,
		63, 1, 10, 1, 10, 0, 0, 11, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 1, 0, 2, 3, 0, 48, 57, 65, 90, 97, 122, 3,
		0, 9, 10, 13, 13, 32, 32, 69, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 1, 35, 1, 0, 0, 0, 3, 38, 1, 0, 0, 0, 5, 42, 1, 0, 0, 0,
		7, 44, 1, 0, 0, 0, 9, 46, 1, 0, 0, 0, 11, 48, 1, 0, 0, 0, 13, 50, 1, 0,
		0, 0, 15, 53, 1, 0, 0, 0, 17, 56, 1, 0, 0, 0, 19, 58, 1, 0, 0, 0, 21, 61,
		1, 0, 0, 0, 23, 24, 5, 100, 0, 0, 24, 25, 5, 105, 0, 0, 25, 26, 5, 103,
		0, 0, 26, 27, 5, 114, 0, 0, 27, 28, 5, 97, 0, 0, 28, 29, 5, 112, 0, 0,
		29, 36, 5, 104, 0, 0, 30, 31, 5, 103, 0, 0, 31, 32, 5, 114, 0, 0, 32, 33,
		5, 97, 0, 0, 33, 34, 5, 112, 0, 0, 34, 36, 5, 104, 0, 0, 35, 23, 1, 0,
		0, 0, 35, 30, 1, 0, 0, 0, 36, 2, 1, 0, 0, 0, 37, 39, 7, 0, 0, 0, 38, 37,
		1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0,
		41, 4, 1, 0, 0, 0, 42, 43, 5, 123, 0, 0, 43, 6, 1, 0, 0, 0, 44, 45, 5,
		125, 0, 0, 45, 8, 1, 0, 0, 0, 46, 47, 5, 91, 0, 0, 47, 10, 1, 0, 0, 0,
		48, 49, 5, 93, 0, 0, 49, 12, 1, 0, 0, 0, 50, 51, 5, 45, 0, 0, 51, 52, 5,
		62, 0, 0, 52, 14, 1, 0, 0, 0, 53, 54, 5, 61, 0, 0, 54, 55, 5, 61, 0, 0,
		55, 16, 1, 0, 0, 0, 56, 57, 5, 61, 0, 0, 57, 18, 1, 0, 0, 0, 58, 59, 5,
		59, 0, 0, 59, 20, 1, 0, 0, 0, 60, 62, 7, 1, 0, 0, 61, 60, 1, 0, 0, 0, 62,
		63, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65, 1, 0, 0,
		0, 65, 66, 6, 10, 0, 0, 66, 22, 1, 0, 0, 0, 4, 0, 35, 40, 63, 1, 6, 0,
		0,
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

// lexDotGraphLexerInit initializes any static state used to implement lexDotGraphLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewlexDotGraphLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LexDotGraphLexerInit() {
	staticData := &lexdotgraphlexerLexerStaticData
	staticData.once.Do(lexdotgraphlexerLexerInit)
}

// NewlexDotGraphLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewlexDotGraphLexer(input antlr.CharStream) *lexDotGraphLexer {
	LexDotGraphLexerInit()
	l := new(lexDotGraphLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &lexdotgraphlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "lexDotGraph.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// lexDotGraphLexer tokens.
const (
	lexDotGraphLexerGRAPH_TYPE           = 1
	lexDotGraphLexerNAME                 = 2
	lexDotGraphLexerOPEN_BRACE           = 3
	lexDotGraphLexerCLOSE_BRACE          = 4
	lexDotGraphLexerOPEN_SQUARE_BRACKET  = 5
	lexDotGraphLexerCLOSE_SQUARE_BRACKET = 6
	lexDotGraphLexerDIRECT_EDGE_OP       = 7
	lexDotGraphLexerUNDIRECT_EDGE_OP     = 8
	lexDotGraphLexerEQUALS_OP            = 9
	lexDotGraphLexerSEMICOLON            = 10
	lexDotGraphLexerWS                   = 11
)
