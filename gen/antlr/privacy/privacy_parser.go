// Code generated from privacy.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // privacy

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

type privacyParser struct {
	*antlr.BaseParser
}

var PrivacyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func privacyParserInit() {
	staticData := &PrivacyParserStaticData
	staticData.LiteralNames = []string{
		"", "'ALLOW'", "'DENY'", "'TO'", "'WHEN'", "';'", "'share'", "'collect'",
		"'use'", "'third parties'", "'partners'", "'internal services'", "'AND'",
		"'OR'", "'time_range'", "'='", "'location'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"TIME_RANGE", "LOCATION", "WS",
	}
	staticData.RuleNames = []string{
		"policy", "rule", "action", "target", "condition", "condition_",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 48, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 4, 0, 14, 8, 0, 11, 0, 12, 0, 15, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 24, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 5, 4, 35, 8, 4, 10, 4, 12, 4, 38, 9, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 3, 5, 46, 8, 5, 1, 5, 0, 0, 6, 0, 2, 4, 6, 8, 10, 0,
		4, 1, 0, 1, 2, 1, 0, 6, 8, 1, 0, 9, 11, 1, 0, 12, 13, 45, 0, 13, 1, 0,
		0, 0, 2, 17, 1, 0, 0, 0, 4, 27, 1, 0, 0, 0, 6, 29, 1, 0, 0, 0, 8, 31, 1,
		0, 0, 0, 10, 45, 1, 0, 0, 0, 12, 14, 3, 2, 1, 0, 13, 12, 1, 0, 0, 0, 14,
		15, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 1, 1, 0, 0,
		0, 17, 18, 7, 0, 0, 0, 18, 19, 3, 4, 2, 0, 19, 20, 5, 3, 0, 0, 20, 23,
		3, 6, 3, 0, 21, 22, 5, 4, 0, 0, 22, 24, 3, 8, 4, 0, 23, 21, 1, 0, 0, 0,
		23, 24, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26, 5, 5, 0, 0, 26, 3, 1, 0,
		0, 0, 27, 28, 7, 1, 0, 0, 28, 5, 1, 0, 0, 0, 29, 30, 7, 2, 0, 0, 30, 7,
		1, 0, 0, 0, 31, 36, 3, 10, 5, 0, 32, 33, 7, 3, 0, 0, 33, 35, 3, 10, 5,
		0, 34, 32, 1, 0, 0, 0, 35, 38, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37,
		1, 0, 0, 0, 37, 9, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 39, 40, 5, 14, 0, 0,
		40, 41, 5, 15, 0, 0, 41, 46, 5, 17, 0, 0, 42, 43, 5, 16, 0, 0, 43, 44,
		5, 15, 0, 0, 44, 46, 5, 18, 0, 0, 45, 39, 1, 0, 0, 0, 45, 42, 1, 0, 0,
		0, 46, 11, 1, 0, 0, 0, 4, 15, 23, 36, 45,
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

// privacyParserInit initializes any static state used to implement privacyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewprivacyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PrivacyParserInit() {
	staticData := &PrivacyParserStaticData
	staticData.once.Do(privacyParserInit)
}

// NewprivacyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewprivacyParser(input antlr.TokenStream) *privacyParser {
	PrivacyParserInit()
	this := new(privacyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PrivacyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "privacy.g4"

	return this
}

// privacyParser tokens.
const (
	privacyParserEOF        = antlr.TokenEOF
	privacyParserT__0       = 1
	privacyParserT__1       = 2
	privacyParserT__2       = 3
	privacyParserT__3       = 4
	privacyParserT__4       = 5
	privacyParserT__5       = 6
	privacyParserT__6       = 7
	privacyParserT__7       = 8
	privacyParserT__8       = 9
	privacyParserT__9       = 10
	privacyParserT__10      = 11
	privacyParserT__11      = 12
	privacyParserT__12      = 13
	privacyParserT__13      = 14
	privacyParserT__14      = 15
	privacyParserT__15      = 16
	privacyParserTIME_RANGE = 17
	privacyParserLOCATION   = 18
	privacyParserWS         = 19
)

// privacyParser rules.
const (
	privacyParserRULE_policy     = 0
	privacyParserRULE_rule       = 1
	privacyParserRULE_action     = 2
	privacyParserRULE_target     = 3
	privacyParserRULE_condition  = 4
	privacyParserRULE_condition_ = 5
)

// IPolicyContext is an interface to support dynamic dispatch.
type IPolicyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRule_() []IRuleContext
	Rule_(i int) IRuleContext

	// IsPolicyContext differentiates from other interfaces.
	IsPolicyContext()
}

type PolicyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPolicyContext() *PolicyContext {
	var p = new(PolicyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = privacyParserRULE_policy
	return p
}

func InitEmptyPolicyContext(p *PolicyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = privacyParserRULE_policy
}

func (*PolicyContext) IsPolicyContext() {}

func NewPolicyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolicyContext {
	var p = new(PolicyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = privacyParserRULE_policy

	return p
}

func (s *PolicyContext) GetParser() antlr.Parser { return s.parser }

func (s *PolicyContext) AllRule_() []IRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleContext); ok {
			len++
		}
	}

	tst := make([]IRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleContext); ok {
			tst[i] = t.(IRuleContext)
			i++
		}
	}

	return tst
}

func (s *PolicyContext) Rule_(i int) IRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleContext); ok {
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

	return t.(IRuleContext)
}

func (s *PolicyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolicyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolicyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(privacyListener); ok {
		listenerT.EnterPolicy(s)
	}
}

func (s *PolicyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(privacyListener); ok {
		listenerT.ExitPolicy(s)
	}
}

func (p *privacyParser) Policy() (localctx IPolicyContext) {
	localctx = NewPolicyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, privacyParserRULE_policy)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == privacyParserT__0 || _la == privacyParserT__1 {
		{
			p.SetState(12)
			p.Rule_()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IRuleContext is an interface to support dynamic dispatch.
type IRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Action_() IActionContext
	Target() ITargetContext
	Condition() IConditionContext

	// IsRuleContext differentiates from other interfaces.
	IsRuleContext()
}

type RuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleContext() *RuleContext {
	var p = new(RuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = privacyParserRULE_rule
	return p
}

func InitEmptyRuleContext(p *RuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = privacyParserRULE_rule
}

func (*RuleContext) IsRuleContext() {}

func NewRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleContext {
	var p = new(RuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = privacyParserRULE_rule

	return p
}

func (s *RuleContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleContext) Action_() IActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionContext)
}

func (s *RuleContext) Target() ITargetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *RuleContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *RuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(privacyListener); ok {
		listenerT.EnterRule(s)
	}
}

func (s *RuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(privacyListener); ok {
		listenerT.ExitRule(s)
	}
}

func (p *privacyParser) Rule_() (localctx IRuleContext) {
	localctx = NewRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, privacyParserRULE_rule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(17)
		_la = p.GetTokenStream().LA(1)

		if !(_la == privacyParserT__0 || _la == privacyParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(18)
		p.Action_()
	}
	{
		p.SetState(19)
		p.Match(privacyParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(20)
		p.Target()
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == privacyParserT__3 {
		{
			p.SetState(21)
			p.Match(privacyParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(22)
			p.Condition()
		}

	}
	{
		p.SetState(25)
		p.Match(privacyParserT__4)
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

// IActionContext is an interface to support dynamic dispatch.
type IActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsActionContext differentiates from other interfaces.
	IsActionContext()
}

type ActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionContext() *ActionContext {
	var p = new(ActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = privacyParserRULE_action
	return p
}

func InitEmptyActionContext(p *ActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = privacyParserRULE_action
}

func (*ActionContext) IsActionContext() {}

func NewActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionContext {
	var p = new(ActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = privacyParserRULE_action

	return p
}

func (s *ActionContext) GetParser() antlr.Parser { return s.parser }
func (s *ActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(privacyListener); ok {
		listenerT.EnterAction(s)
	}
}

func (s *ActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(privacyListener); ok {
		listenerT.ExitAction(s)
	}
}

func (p *privacyParser) Action_() (localctx IActionContext) {
	localctx = NewActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, privacyParserRULE_action)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&448) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = privacyParserRULE_target
	return p
}

func InitEmptyTargetContext(p *TargetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = privacyParserRULE_target
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = privacyParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }
func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(privacyListener); ok {
		listenerT.EnterTarget(s)
	}
}

func (s *TargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(privacyListener); ok {
		listenerT.ExitTarget(s)
	}
}

func (p *privacyParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, privacyParserRULE_target)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3584) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCondition_() []ICondition_Context
	Condition_(i int) ICondition_Context

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = privacyParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = privacyParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = privacyParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) AllCondition_() []ICondition_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICondition_Context); ok {
			len++
		}
	}

	tst := make([]ICondition_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICondition_Context); ok {
			tst[i] = t.(ICondition_Context)
			i++
		}
	}

	return tst
}

func (s *ConditionContext) Condition_(i int) ICondition_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondition_Context); ok {
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

	return t.(ICondition_Context)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(privacyListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(privacyListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *privacyParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, privacyParserRULE_condition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Condition_()
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == privacyParserT__11 || _la == privacyParserT__12 {
		{
			p.SetState(32)
			_la = p.GetTokenStream().LA(1)

			if !(_la == privacyParserT__11 || _la == privacyParserT__12) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(33)
			p.Condition_()
		}

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ICondition_Context is an interface to support dynamic dispatch.
type ICondition_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TIME_RANGE() antlr.TerminalNode
	LOCATION() antlr.TerminalNode

	// IsCondition_Context differentiates from other interfaces.
	IsCondition_Context()
}

type Condition_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondition_Context() *Condition_Context {
	var p = new(Condition_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = privacyParserRULE_condition_
	return p
}

func InitEmptyCondition_Context(p *Condition_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = privacyParserRULE_condition_
}

func (*Condition_Context) IsCondition_Context() {}

func NewCondition_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Condition_Context {
	var p = new(Condition_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = privacyParserRULE_condition_

	return p
}

func (s *Condition_Context) GetParser() antlr.Parser { return s.parser }

func (s *Condition_Context) TIME_RANGE() antlr.TerminalNode {
	return s.GetToken(privacyParserTIME_RANGE, 0)
}

func (s *Condition_Context) LOCATION() antlr.TerminalNode {
	return s.GetToken(privacyParserLOCATION, 0)
}

func (s *Condition_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Condition_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Condition_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(privacyListener); ok {
		listenerT.EnterCondition_(s)
	}
}

func (s *Condition_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(privacyListener); ok {
		listenerT.ExitCondition_(s)
	}
}

func (p *privacyParser) Condition_() (localctx ICondition_Context) {
	localctx = NewCondition_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, privacyParserRULE_condition_)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case privacyParserT__13:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Match(privacyParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Match(privacyParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Match(privacyParserTIME_RANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case privacyParserT__15:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Match(privacyParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.Match(privacyParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Match(privacyParserLOCATION)
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
