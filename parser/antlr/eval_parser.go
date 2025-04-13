// Code generated from eval.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // eval

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type evalParser struct {
	*antlr.BaseParser
}

var EvalParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func evalParserInit() {
	staticData := &EvalParserStaticData
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
		"program", "variable", "statement", "function",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 69, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 17, 8, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 28, 8, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 5, 2, 48, 8, 2, 10, 2, 12, 2, 51, 9, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 5, 3, 58, 8, 3, 10, 3, 12, 3, 61, 9, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 3, 3, 67, 8, 3, 1, 3, 0, 1, 4, 4, 0, 2, 4, 6, 0, 5, 2, 0, 9, 9, 11,
		11, 1, 0, 5, 7, 1, 0, 8, 9, 1, 0, 13, 16, 1, 0, 17, 18, 79, 0, 8, 1, 0,
		0, 0, 2, 16, 1, 0, 0, 0, 4, 27, 1, 0, 0, 0, 6, 66, 1, 0, 0, 0, 8, 9, 3,
		4, 2, 0, 9, 10, 5, 0, 0, 1, 10, 1, 1, 0, 0, 0, 11, 17, 5, 19, 0, 0, 12,
		17, 5, 22, 0, 0, 13, 17, 5, 24, 0, 0, 14, 17, 5, 23, 0, 0, 15, 17, 5, 25,
		0, 0, 16, 11, 1, 0, 0, 0, 16, 12, 1, 0, 0, 0, 16, 13, 1, 0, 0, 0, 16, 14,
		1, 0, 0, 0, 16, 15, 1, 0, 0, 0, 17, 3, 1, 0, 0, 0, 18, 19, 6, 2, -1, 0,
		19, 28, 3, 2, 1, 0, 20, 28, 3, 6, 3, 0, 21, 22, 7, 0, 0, 0, 22, 28, 3,
		4, 2, 8, 23, 24, 5, 1, 0, 0, 24, 25, 3, 4, 2, 0, 25, 26, 5, 2, 0, 0, 26,
		28, 1, 0, 0, 0, 27, 18, 1, 0, 0, 0, 27, 20, 1, 0, 0, 0, 27, 21, 1, 0, 0,
		0, 27, 23, 1, 0, 0, 0, 28, 49, 1, 0, 0, 0, 29, 30, 10, 7, 0, 0, 30, 31,
		7, 1, 0, 0, 31, 48, 3, 4, 2, 8, 32, 33, 10, 6, 0, 0, 33, 34, 7, 2, 0, 0,
		34, 48, 3, 4, 2, 7, 35, 36, 10, 5, 0, 0, 36, 37, 7, 3, 0, 0, 37, 48, 3,
		4, 2, 6, 38, 39, 10, 4, 0, 0, 39, 40, 7, 4, 0, 0, 40, 48, 3, 4, 2, 5, 41,
		42, 10, 3, 0, 0, 42, 43, 5, 10, 0, 0, 43, 48, 3, 4, 2, 4, 44, 45, 10, 2,
		0, 0, 45, 46, 5, 12, 0, 0, 46, 48, 3, 4, 2, 3, 47, 29, 1, 0, 0, 0, 47,
		32, 1, 0, 0, 0, 47, 35, 1, 0, 0, 0, 47, 38, 1, 0, 0, 0, 47, 41, 1, 0, 0,
		0, 47, 44, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 50,
		1, 0, 0, 0, 50, 5, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 53, 5, 25, 0, 0,
		53, 54, 5, 1, 0, 0, 54, 59, 3, 4, 2, 0, 55, 56, 5, 3, 0, 0, 56, 58, 3,
		4, 2, 0, 57, 55, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59,
		60, 1, 0, 0, 0, 60, 62, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 63, 5, 2, 0,
		0, 63, 67, 1, 0, 0, 0, 64, 65, 5, 25, 0, 0, 65, 67, 5, 4, 0, 0, 66, 52,
		1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 7, 1, 0, 0, 0, 6, 16, 27, 47, 49, 59,
		66,
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

// evalParserInit initializes any static state used to implement evalParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewevalParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EvalParserInit() {
	staticData := &EvalParserStaticData
	staticData.once.Do(evalParserInit)
}

// NewevalParser produces a new parser instance for the optional input antlr.TokenStream.
func NewevalParser(input antlr.TokenStream) *evalParser {
	EvalParserInit()
	this := new(evalParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &EvalParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "eval.g4"

	return this
}

// evalParser tokens.
const (
	evalParserEOF        = antlr.TokenEOF
	evalParserT__0       = 1
	evalParserT__1       = 2
	evalParserT__2       = 3
	evalParserT__3       = 4
	evalParserMUL        = 5
	evalParserDIV        = 6
	evalParserMOD        = 7
	evalParserADD        = 8
	evalParserSUB        = 9
	evalParserAND        = 10
	evalParserNOT        = 11
	evalParserOR         = 12
	evalParserLE         = 13
	evalParserLT         = 14
	evalParserGE         = 15
	evalParserGT         = 16
	evalParserNE         = 17
	evalParserEQ         = 18
	evalParserBOOL       = 19
	evalParserTRUE       = 20
	evalParserFALSE      = 21
	evalParserSTRING     = 22
	evalParserINTEGER    = 23
	evalParserFLOAT      = 24
	evalParserVAR        = 25
	evalParserWHITESPACE = 26
)

// evalParser rules.
const (
	evalParserRULE_program   = 0
	evalParserRULE_variable  = 1
	evalParserRULE_statement = 2
	evalParserRULE_function  = 3
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statement() IStatementContext
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = evalParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = evalParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = evalParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(evalParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *evalParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, evalParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.statement(0)
	}
	{
		p.SetState(9)
		p.Match(evalParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = evalParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = evalParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = evalParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) CopyAll(ctx *VariableContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IntegerContext struct {
	VariableContext
}

func NewIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerContext {
	var p = new(IntegerContext)

	InitEmptyVariableContext(&p.VariableContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableContext))

	return p
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(evalParserINTEGER, 0)
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatContext struct {
	VariableContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	InitEmptyVariableContext(&p.VariableContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(evalParserFLOAT, 0)
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolContext struct {
	VariableContext
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	InitEmptyVariableContext(&p.VariableContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableContext))

	return p
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(evalParserBOOL, 0)
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarContext struct {
	VariableContext
}

func NewVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarContext {
	var p = new(VarContext)

	InitEmptyVariableContext(&p.VariableContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableContext))

	return p
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) VAR() antlr.TerminalNode {
	return s.GetToken(evalParserVAR, 0)
}

func (s *VarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	VariableContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyVariableContext(&p.VariableContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(evalParserSTRING, 0)
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *evalParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, evalParserRULE_variable)
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case evalParserBOOL:
		localctx = NewBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(11)
			p.Match(evalParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case evalParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(12)
			p.Match(evalParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case evalParserFLOAT:
		localctx = NewFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(13)
			p.Match(evalParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case evalParserINTEGER:
		localctx = NewIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(14)
			p.Match(evalParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case evalParserVAR:
		localctx = NewVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(15)
			p.Match(evalParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = evalParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = evalParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = evalParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OrContext struct {
	StatementContext
	op antlr.Token
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *OrContext) GetOp() antlr.Token { return s.op }

func (s *OrContext) SetOp(v antlr.Token) { s.op = v }

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *OrContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *OrContext) OR() antlr.TerminalNode {
	return s.GetToken(evalParserOR, 0)
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Comparison1Context struct {
	StatementContext
	op antlr.Token
}

func NewComparison1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Comparison1Context {
	var p = new(Comparison1Context)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *Comparison1Context) GetOp() antlr.Token { return s.op }

func (s *Comparison1Context) SetOp(v antlr.Token) { s.op = v }

func (s *Comparison1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparison1Context) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Comparison1Context) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *Comparison1Context) LE() antlr.TerminalNode {
	return s.GetToken(evalParserLE, 0)
}

func (s *Comparison1Context) LT() antlr.TerminalNode {
	return s.GetToken(evalParserLT, 0)
}

func (s *Comparison1Context) GE() antlr.TerminalNode {
	return s.GetToken(evalParserGE, 0)
}

func (s *Comparison1Context) GT() antlr.TerminalNode {
	return s.GetToken(evalParserGT, 0)
}

func (s *Comparison1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitComparison1(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivContext struct {
	StatementContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(evalParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(evalParserDIV, 0)
}

func (s *MulDivContext) MOD() antlr.TerminalNode {
	return s.GetToken(evalParserMOD, 0)
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSUbContext struct {
	StatementContext
	op antlr.Token
}

func NewAddSUbContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSUbContext {
	var p = new(AddSUbContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *AddSUbContext) GetOp() antlr.Token { return s.op }

func (s *AddSUbContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSUbContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSUbContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *AddSUbContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *AddSUbContext) ADD() antlr.TerminalNode {
	return s.GetToken(evalParserADD, 0)
}

func (s *AddSUbContext) SUB() antlr.TerminalNode {
	return s.GetToken(evalParserSUB, 0)
}

func (s *AddSUbContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitAddSUb(s)

	default:
		return t.VisitChildren(s)
	}
}

type Comparison2Context struct {
	StatementContext
	op antlr.Token
}

func NewComparison2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Comparison2Context {
	var p = new(Comparison2Context)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *Comparison2Context) GetOp() antlr.Token { return s.op }

func (s *Comparison2Context) SetOp(v antlr.Token) { s.op = v }

func (s *Comparison2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparison2Context) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Comparison2Context) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *Comparison2Context) NE() antlr.TerminalNode {
	return s.GetToken(evalParserNE, 0)
}

func (s *Comparison2Context) EQ() antlr.TerminalNode {
	return s.GetToken(evalParserEQ, 0)
}

func (s *Comparison2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitComparison2(s)

	default:
		return t.VisitChildren(s)
	}
}

type VContext struct {
	StatementContext
}

func NewVContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VContext {
	var p = new(VContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *VContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitV(s)

	default:
		return t.VisitChildren(s)
	}
}

type FContext struct {
	StatementContext
}

func NewFContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FContext {
	var p = new(FContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *FContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitF(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensContext struct {
	StatementContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndContext struct {
	StatementContext
	op antlr.Token
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *AndContext) GetOp() antlr.Token { return s.op }

func (s *AndContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *AndContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *AndContext) AND() antlr.TerminalNode {
	return s.GetToken(evalParserAND, 0)
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryContext struct {
	StatementContext
	op antlr.Token
}

func NewUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryContext {
	var p = new(UnaryContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *UnaryContext) GetOp() antlr.Token { return s.op }

func (s *UnaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *UnaryContext) NOT() antlr.TerminalNode {
	return s.GetToken(evalParserNOT, 0)
}

func (s *UnaryContext) SUB() antlr.TerminalNode {
	return s.GetToken(evalParserSUB, 0)
}

func (s *UnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *evalParser) Statement() (localctx IStatementContext) {
	return p.statement(0)
}

func (p *evalParser) statement(_p int) (localctx IStatementContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, evalParserRULE_statement, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(19)
			p.Variable()
		}

	case 2:
		localctx = NewFContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(20)
			p.Function()
		}

	case 3:
		localctx = NewUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(21)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == evalParserSUB || _la == evalParserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(22)
			p.statement(8)
		}

	case 4:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)
			p.Match(evalParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(24)
			p.statement(0)
		}
		{
			p.SetState(25)
			p.Match(evalParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, evalParserRULE_statement)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(30)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&224) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(31)
					p.statement(8)
				}

			case 2:
				localctx = NewAddSUbContext(p, NewStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, evalParserRULE_statement)
				p.SetState(32)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(33)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSUbContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == evalParserADD || _la == evalParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSUbContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(34)
					p.statement(7)
				}

			case 3:
				localctx = NewComparison1Context(p, NewStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, evalParserRULE_statement)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(36)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Comparison1Context).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&122880) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Comparison1Context).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(37)
					p.statement(6)
				}

			case 4:
				localctx = NewComparison2Context(p, NewStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, evalParserRULE_statement)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(39)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Comparison2Context).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == evalParserNE || _la == evalParserEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Comparison2Context).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(40)
					p.statement(5)
				}

			case 5:
				localctx = NewAndContext(p, NewStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, evalParserRULE_statement)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(42)

					var _m = p.Match(evalParserAND)

					localctx.(*AndContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(43)
					p.statement(4)
				}

			case 6:
				localctx = NewOrContext(p, NewStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, evalParserRULE_statement)
				p.SetState(44)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(45)

					var _m = p.Match(evalParserOR)

					localctx.(*OrContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(46)
					p.statement(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetF returns the f token.
	GetF() antlr.Token

	// SetF sets the f token.
	SetF(antlr.Token)

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	VAR() antlr.TerminalNode

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	f      antlr.Token
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = evalParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = evalParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = evalParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) GetF() antlr.Token { return s.f }

func (s *FunctionContext) SetF(v antlr.Token) { s.f = v }

func (s *FunctionContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *FunctionContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *FunctionContext) VAR() antlr.TerminalNode {
	return s.GetToken(evalParserVAR, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case evalVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *evalParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, evalParserRULE_function)
	var _la int

	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)

			var _m = p.Match(evalParserVAR)

			localctx.(*FunctionContext).f = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Match(evalParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.statement(0)
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == evalParserT__2 {
			{
				p.SetState(55)
				p.Match(evalParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(56)
				p.statement(0)
			}

			p.SetState(61)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(62)
			p.Match(evalParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)

			var _m = p.Match(evalParserVAR)

			localctx.(*FunctionContext).f = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Match(evalParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *evalParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *StatementContext = nil
		if localctx != nil {
			t = localctx.(*StatementContext)
		}
		return p.Statement_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *evalParser) Statement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
