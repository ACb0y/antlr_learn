// Code generated from ./ArrayInit.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // ArrayInit

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 22, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 11, 10, 2, 12, 2, 14,
	2, 14, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 20, 10, 3, 3, 3, 2, 2, 4, 2,
	4, 2, 2, 2, 21, 2, 6, 3, 2, 2, 2, 4, 19, 3, 2, 2, 2, 6, 7, 7, 3, 2, 2,
	7, 12, 5, 4, 3, 2, 8, 9, 7, 4, 2, 2, 9, 11, 5, 4, 3, 2, 10, 8, 3, 2, 2,
	2, 11, 14, 3, 2, 2, 2, 12, 10, 3, 2, 2, 2, 12, 13, 3, 2, 2, 2, 13, 15,
	3, 2, 2, 2, 14, 12, 3, 2, 2, 2, 15, 16, 7, 5, 2, 2, 16, 3, 3, 2, 2, 2,
	17, 20, 5, 2, 2, 2, 18, 20, 7, 6, 2, 2, 19, 17, 3, 2, 2, 2, 19, 18, 3,
	2, 2, 2, 20, 5, 3, 2, 2, 2, 4, 12, 19,
}
var literalNames = []string{
	"", "'{'", "','", "'}'",
}
var symbolicNames = []string{
	"", "", "", "", "INT", "WS",
}

var ruleNames = []string{
	"init", "value",
}

type ArrayInitParser struct {
	*antlr.BaseParser
}

// NewArrayInitParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ArrayInitParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewArrayInitParser(input antlr.TokenStream) *ArrayInitParser {
	this := new(ArrayInitParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ArrayInit.g4"

	return this
}

// ArrayInitParser tokens.
const (
	ArrayInitParserEOF  = antlr.TokenEOF
	ArrayInitParserT__0 = 1
	ArrayInitParserT__1 = 2
	ArrayInitParserT__2 = 3
	ArrayInitParserINT  = 4
	ArrayInitParserWS   = 5
)

// ArrayInitParser rules.
const (
	ArrayInitParserRULE_init  = 0
	ArrayInitParserRULE_value = 1
)

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArrayInitParserRULE_init
	return p
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArrayInitParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *InitContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayInitListener); ok {
		listenerT.EnterInit(s)
	}
}

func (s *InitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayInitListener); ok {
		listenerT.ExitInit(s)
	}
}

func (p *ArrayInitParser) Init() (localctx IInitContext) {
	this := p
	_ = this

	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ArrayInitParserRULE_init)
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
		p.SetState(4)
		p.Match(ArrayInitParserT__0)
	}
	{
		p.SetState(5)
		p.Value()
	}
	p.SetState(10)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ArrayInitParserT__1 {
		{
			p.SetState(6)
			p.Match(ArrayInitParserT__1)
		}
		{
			p.SetState(7)
			p.Value()
		}

		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(13)
		p.Match(ArrayInitParserT__2)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArrayInitParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArrayInitParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Init() IInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitContext)
}

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(ArrayInitParserINT, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayInitListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayInitListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *ArrayInitParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ArrayInitParserRULE_value)

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

	p.SetState(17)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ArrayInitParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(15)
			p.Init()
		}

	case ArrayInitParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(16)
			p.Match(ArrayInitParserINT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
