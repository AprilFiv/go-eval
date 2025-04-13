// Code generated from eval.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type evalLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var EvalLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func evallexerLexerInit() {
	staticData := &EvalLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "','", "'()'", "'*'", "'/'", "'%'", "'+'", "'-'",
		"'&&'", "'!'", "'||'", "'<='", "'<'", "'>='", "'>'", "", "'=='", "",
		"'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "MUL", "DIV", "MOD", "ADD", "SUB", "AND", "NOT",
		"OR", "LE", "LT", "GE", "GT", "NE", "EQ", "BOOL", "TRUE", "FALSE", "STRING",
		"INTEGER", "FLOAT", "VAR", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "MUL", "DIV", "MOD", "ADD", "SUB", "AND",
		"NOT", "OR", "LE", "LT", "GE", "GT", "NE", "EQ", "BOOL", "TRUE", "FALSE",
		"STRING", "INTEGER", "FLOAT", "VAR", "WHITESPACE", "DIGIT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 166, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 97,
		8, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 3, 18, 104, 8, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21,
		1, 21, 5, 21, 119, 8, 21, 10, 21, 12, 21, 122, 9, 21, 1, 21, 1, 21, 1,
		21, 5, 21, 127, 8, 21, 10, 21, 12, 21, 130, 9, 21, 1, 21, 3, 21, 133, 8,
		21, 1, 22, 4, 22, 136, 8, 22, 11, 22, 12, 22, 137, 1, 23, 4, 23, 141, 8,
		23, 11, 23, 12, 23, 142, 1, 23, 1, 23, 4, 23, 147, 8, 23, 11, 23, 12, 23,
		148, 1, 24, 1, 24, 5, 24, 153, 8, 24, 10, 24, 12, 24, 156, 9, 24, 1, 25,
		4, 25, 159, 8, 25, 11, 25, 12, 25, 160, 1, 25, 1, 25, 1, 26, 1, 26, 2,
		120, 128, 0, 27, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53,
		0, 1, 0, 4, 3, 0, 65, 90, 95, 95, 97, 122, 5, 0, 46, 46, 48, 57, 65, 90,
		95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 174, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 1, 55, 1, 0, 0, 0, 3, 57,
		1, 0, 0, 0, 5, 59, 1, 0, 0, 0, 7, 61, 1, 0, 0, 0, 9, 64, 1, 0, 0, 0, 11,
		66, 1, 0, 0, 0, 13, 68, 1, 0, 0, 0, 15, 70, 1, 0, 0, 0, 17, 72, 1, 0, 0,
		0, 19, 74, 1, 0, 0, 0, 21, 77, 1, 0, 0, 0, 23, 79, 1, 0, 0, 0, 25, 82,
		1, 0, 0, 0, 27, 85, 1, 0, 0, 0, 29, 87, 1, 0, 0, 0, 31, 90, 1, 0, 0, 0,
		33, 96, 1, 0, 0, 0, 35, 98, 1, 0, 0, 0, 37, 103, 1, 0, 0, 0, 39, 105, 1,
		0, 0, 0, 41, 110, 1, 0, 0, 0, 43, 132, 1, 0, 0, 0, 45, 135, 1, 0, 0, 0,
		47, 140, 1, 0, 0, 0, 49, 150, 1, 0, 0, 0, 51, 158, 1, 0, 0, 0, 53, 164,
		1, 0, 0, 0, 55, 56, 5, 40, 0, 0, 56, 2, 1, 0, 0, 0, 57, 58, 5, 41, 0, 0,
		58, 4, 1, 0, 0, 0, 59, 60, 5, 44, 0, 0, 60, 6, 1, 0, 0, 0, 61, 62, 5, 40,
		0, 0, 62, 63, 5, 41, 0, 0, 63, 8, 1, 0, 0, 0, 64, 65, 5, 42, 0, 0, 65,
		10, 1, 0, 0, 0, 66, 67, 5, 47, 0, 0, 67, 12, 1, 0, 0, 0, 68, 69, 5, 37,
		0, 0, 69, 14, 1, 0, 0, 0, 70, 71, 5, 43, 0, 0, 71, 16, 1, 0, 0, 0, 72,
		73, 5, 45, 0, 0, 73, 18, 1, 0, 0, 0, 74, 75, 5, 38, 0, 0, 75, 76, 5, 38,
		0, 0, 76, 20, 1, 0, 0, 0, 77, 78, 5, 33, 0, 0, 78, 22, 1, 0, 0, 0, 79,
		80, 5, 124, 0, 0, 80, 81, 5, 124, 0, 0, 81, 24, 1, 0, 0, 0, 82, 83, 5,
		60, 0, 0, 83, 84, 5, 61, 0, 0, 84, 26, 1, 0, 0, 0, 85, 86, 5, 60, 0, 0,
		86, 28, 1, 0, 0, 0, 87, 88, 5, 62, 0, 0, 88, 89, 5, 61, 0, 0, 89, 30, 1,
		0, 0, 0, 90, 91, 5, 62, 0, 0, 91, 32, 1, 0, 0, 0, 92, 93, 5, 126, 0, 0,
		93, 97, 5, 61, 0, 0, 94, 95, 5, 33, 0, 0, 95, 97, 5, 61, 0, 0, 96, 92,
		1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 34, 1, 0, 0, 0, 98, 99, 5, 61, 0, 0,
		99, 100, 5, 61, 0, 0, 100, 36, 1, 0, 0, 0, 101, 104, 3, 39, 19, 0, 102,
		104, 3, 41, 20, 0, 103, 101, 1, 0, 0, 0, 103, 102, 1, 0, 0, 0, 104, 38,
		1, 0, 0, 0, 105, 106, 5, 116, 0, 0, 106, 107, 5, 114, 0, 0, 107, 108, 5,
		117, 0, 0, 108, 109, 5, 101, 0, 0, 109, 40, 1, 0, 0, 0, 110, 111, 5, 102,
		0, 0, 111, 112, 5, 97, 0, 0, 112, 113, 5, 108, 0, 0, 113, 114, 5, 115,
		0, 0, 114, 115, 5, 101, 0, 0, 115, 42, 1, 0, 0, 0, 116, 120, 5, 34, 0,
		0, 117, 119, 9, 0, 0, 0, 118, 117, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120,
		121, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122, 120,
		1, 0, 0, 0, 123, 133, 5, 34, 0, 0, 124, 128, 5, 39, 0, 0, 125, 127, 9,
		0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 129, 1, 0, 0,
		0, 128, 126, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131,
		133, 5, 39, 0, 0, 132, 116, 1, 0, 0, 0, 132, 124, 1, 0, 0, 0, 133, 44,
		1, 0, 0, 0, 134, 136, 3, 53, 26, 0, 135, 134, 1, 0, 0, 0, 136, 137, 1,
		0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 46, 1, 0, 0,
		0, 139, 141, 3, 53, 26, 0, 140, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0,
		142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144,
		146, 5, 46, 0, 0, 145, 147, 3, 53, 26, 0, 146, 145, 1, 0, 0, 0, 147, 148,
		1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 48, 1, 0,
		0, 0, 150, 154, 7, 0, 0, 0, 151, 153, 7, 1, 0, 0, 152, 151, 1, 0, 0, 0,
		153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155,
		50, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 159, 7, 2, 0, 0, 158, 157, 1,
		0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0,
		0, 161, 162, 1, 0, 0, 0, 162, 163, 6, 25, 0, 0, 163, 52, 1, 0, 0, 0, 164,
		165, 7, 3, 0, 0, 165, 54, 1, 0, 0, 0, 11, 0, 96, 103, 120, 128, 132, 137,
		142, 148, 154, 160, 1, 6, 0, 0,
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

// evalLexerInit initializes any static state used to implement evalLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewevalLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func EvalLexerInit() {
	staticData := &EvalLexerLexerStaticData
	staticData.once.Do(evallexerLexerInit)
}

// NewevalLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewevalLexer(input antlr.CharStream) *evalLexer {
	EvalLexerInit()
	l := new(evalLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &EvalLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "eval.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// evalLexer tokens.
const (
	evalLexerT__0       = 1
	evalLexerT__1       = 2
	evalLexerT__2       = 3
	evalLexerT__3       = 4
	evalLexerMUL        = 5
	evalLexerDIV        = 6
	evalLexerMOD        = 7
	evalLexerADD        = 8
	evalLexerSUB        = 9
	evalLexerAND        = 10
	evalLexerNOT        = 11
	evalLexerOR         = 12
	evalLexerLE         = 13
	evalLexerLT         = 14
	evalLexerGE         = 15
	evalLexerGT         = 16
	evalLexerNE         = 17
	evalLexerEQ         = 18
	evalLexerBOOL       = 19
	evalLexerTRUE       = 20
	evalLexerFALSE      = 21
	evalLexerSTRING     = 22
	evalLexerINTEGER    = 23
	evalLexerFLOAT      = 24
	evalLexerVAR        = 25
	evalLexerWHITESPACE = 26
)
