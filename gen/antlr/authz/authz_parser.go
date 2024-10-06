// Code generated from authz.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // authz

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

type authzParser struct {
	*antlr.BaseParser
}

var AuthzParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func authzParserInit() {
	staticData := &AuthzParserStaticData
	staticData.LiteralNames = []string{
		"", "'ALLOW'", "'DENY'", "'Exec'", "'TO'", "'WHEN'", "';'", "'user'",
		"'service'", "'share'", "'collect'", "'use'", "'data'", "'feature'",
		"'AND'", "'OR'", "'time_range'", "'='", "'location'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "TIME_RANGE", "LOCATION", "WS",
	}
	staticData.RuleNames = []string{
		"policy", "rule", "subject", "action", "object", "condition", "condition_",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 54, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 4, 0, 16, 8, 0, 11, 0, 12, 0, 17, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 41, 8, 5, 10,
		5, 12, 5, 44, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 52, 8, 6,
		1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 5, 1, 0, 1, 2, 1, 0, 7, 8, 1,
		0, 9, 11, 2, 0, 8, 8, 12, 13, 1, 0, 14, 15, 50, 0, 15, 1, 0, 0, 0, 2, 19,
		1, 0, 0, 0, 4, 31, 1, 0, 0, 0, 6, 33, 1, 0, 0, 0, 8, 35, 1, 0, 0, 0, 10,
		37, 1, 0, 0, 0, 12, 51, 1, 0, 0, 0, 14, 16, 3, 2, 1, 0, 15, 14, 1, 0, 0,
		0, 16, 17, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 1, 1,
		0, 0, 0, 19, 20, 7, 0, 0, 0, 20, 21, 3, 4, 2, 0, 21, 22, 5, 3, 0, 0, 22,
		23, 3, 6, 3, 0, 23, 24, 5, 4, 0, 0, 24, 27, 3, 8, 4, 0, 25, 26, 5, 5, 0,
		0, 26, 28, 3, 10, 5, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 29,
		1, 0, 0, 0, 29, 30, 5, 6, 0, 0, 30, 3, 1, 0, 0, 0, 31, 32, 7, 1, 0, 0,
		32, 5, 1, 0, 0, 0, 33, 34, 7, 2, 0, 0, 34, 7, 1, 0, 0, 0, 35, 36, 7, 3,
		0, 0, 36, 9, 1, 0, 0, 0, 37, 42, 3, 12, 6, 0, 38, 39, 7, 4, 0, 0, 39, 41,
		3, 12, 6, 0, 40, 38, 1, 0, 0, 0, 41, 44, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0,
		42, 43, 1, 0, 0, 0, 43, 11, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 45, 46, 5,
		16, 0, 0, 46, 47, 5, 17, 0, 0, 47, 52, 5, 19, 0, 0, 48, 49, 5, 18, 0, 0,
		49, 50, 5, 17, 0, 0, 50, 52, 5, 20, 0, 0, 51, 45, 1, 0, 0, 0, 51, 48, 1,
		0, 0, 0, 52, 13, 1, 0, 0, 0, 4, 17, 27, 42, 51,
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

// authzParserInit initializes any static state used to implement authzParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewauthzParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AuthzParserInit() {
	staticData := &AuthzParserStaticData
	staticData.once.Do(authzParserInit)
}

// NewauthzParser produces a new parser instance for the optional input antlr.TokenStream.
func NewauthzParser(input antlr.TokenStream) *authzParser {
	AuthzParserInit()
	this := new(authzParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AuthzParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "authz.g4"

	return this
}

// authzParser tokens.
const (
	authzParserEOF        = antlr.TokenEOF
	authzParserT__0       = 1
	authzParserT__1       = 2
	authzParserT__2       = 3
	authzParserT__3       = 4
	authzParserT__4       = 5
	authzParserT__5       = 6
	authzParserT__6       = 7
	authzParserT__7       = 8
	authzParserT__8       = 9
	authzParserT__9       = 10
	authzParserT__10      = 11
	authzParserT__11      = 12
	authzParserT__12      = 13
	authzParserT__13      = 14
	authzParserT__14      = 15
	authzParserT__15      = 16
	authzParserT__16      = 17
	authzParserT__17      = 18
	authzParserTIME_RANGE = 19
	authzParserLOCATION   = 20
	authzParserWS         = 21
)

// authzParser rules.
const (
	authzParserRULE_policy     = 0
	authzParserRULE_rule       = 1
	authzParserRULE_subject    = 2
	authzParserRULE_action     = 3
	authzParserRULE_object     = 4
	authzParserRULE_condition  = 5
	authzParserRULE_condition_ = 6
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
	p.RuleIndex = authzParserRULE_policy
	return p
}

func InitEmptyPolicyContext(p *PolicyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = authzParserRULE_policy
}

func (*PolicyContext) IsPolicyContext() {}

func NewPolicyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolicyContext {
	var p = new(PolicyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = authzParserRULE_policy

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
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.EnterPolicy(s)
	}
}

func (s *PolicyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.ExitPolicy(s)
	}
}

func (p *authzParser) Policy() (localctx IPolicyContext) {
	localctx = NewPolicyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, authzParserRULE_policy)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == authzParserT__0 || _la == authzParserT__1 {
		{
			p.SetState(14)
			p.Rule_()
		}

		p.SetState(17)
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
	Subject() ISubjectContext
	Action_() IActionContext
	Object() IObjectContext
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
	p.RuleIndex = authzParserRULE_rule
	return p
}

func InitEmptyRuleContext(p *RuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = authzParserRULE_rule
}

func (*RuleContext) IsRuleContext() {}

func NewRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleContext {
	var p = new(RuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = authzParserRULE_rule

	return p
}

func (s *RuleContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleContext) Subject() ISubjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubjectContext)
}

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

func (s *RuleContext) Object() IObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
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
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.EnterRule(s)
	}
}

func (s *RuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.ExitRule(s)
	}
}

func (p *authzParser) Rule_() (localctx IRuleContext) {
	localctx = NewRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, authzParserRULE_rule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		_la = p.GetTokenStream().LA(1)

		if !(_la == authzParserT__0 || _la == authzParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(20)
		p.Subject()
	}
	{
		p.SetState(21)
		p.Match(authzParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)
		p.Action_()
	}
	{
		p.SetState(23)
		p.Match(authzParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(24)
		p.Object()
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == authzParserT__4 {
		{
			p.SetState(25)
			p.Match(authzParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
			p.Condition()
		}

	}
	{
		p.SetState(29)
		p.Match(authzParserT__5)
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

// ISubjectContext is an interface to support dynamic dispatch.
type ISubjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSubjectContext differentiates from other interfaces.
	IsSubjectContext()
}

type SubjectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubjectContext() *SubjectContext {
	var p = new(SubjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = authzParserRULE_subject
	return p
}

func InitEmptySubjectContext(p *SubjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = authzParserRULE_subject
}

func (*SubjectContext) IsSubjectContext() {}

func NewSubjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubjectContext {
	var p = new(SubjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = authzParserRULE_subject

	return p
}

func (s *SubjectContext) GetParser() antlr.Parser { return s.parser }
func (s *SubjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.EnterSubject(s)
	}
}

func (s *SubjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.ExitSubject(s)
	}
}

func (p *authzParser) Subject() (localctx ISubjectContext) {
	localctx = NewSubjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, authzParserRULE_subject)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		_la = p.GetTokenStream().LA(1)

		if !(_la == authzParserT__6 || _la == authzParserT__7) {
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
	p.RuleIndex = authzParserRULE_action
	return p
}

func InitEmptyActionContext(p *ActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = authzParserRULE_action
}

func (*ActionContext) IsActionContext() {}

func NewActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionContext {
	var p = new(ActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = authzParserRULE_action

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
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.EnterAction(s)
	}
}

func (s *ActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.ExitAction(s)
	}
}

func (p *authzParser) Action_() (localctx IActionContext) {
	localctx = NewActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, authzParserRULE_action)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
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

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = authzParserRULE_object
	return p
}

func InitEmptyObjectContext(p *ObjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = authzParserRULE_object
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = authzParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }
func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.ExitObject(s)
	}
}

func (p *authzParser) Object() (localctx IObjectContext) {
	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, authzParserRULE_object)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&12544) != 0) {
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
	p.RuleIndex = authzParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = authzParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = authzParserRULE_condition

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
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *authzParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, authzParserRULE_condition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Condition_()
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == authzParserT__13 || _la == authzParserT__14 {
		{
			p.SetState(38)
			_la = p.GetTokenStream().LA(1)

			if !(_la == authzParserT__13 || _la == authzParserT__14) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(39)
			p.Condition_()
		}

		p.SetState(44)
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
	p.RuleIndex = authzParserRULE_condition_
	return p
}

func InitEmptyCondition_Context(p *Condition_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = authzParserRULE_condition_
}

func (*Condition_Context) IsCondition_Context() {}

func NewCondition_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Condition_Context {
	var p = new(Condition_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = authzParserRULE_condition_

	return p
}

func (s *Condition_Context) GetParser() antlr.Parser { return s.parser }

func (s *Condition_Context) TIME_RANGE() antlr.TerminalNode {
	return s.GetToken(authzParserTIME_RANGE, 0)
}

func (s *Condition_Context) LOCATION() antlr.TerminalNode {
	return s.GetToken(authzParserLOCATION, 0)
}

func (s *Condition_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Condition_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Condition_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.EnterCondition_(s)
	}
}

func (s *Condition_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(authzListener); ok {
		listenerT.ExitCondition_(s)
	}
}

func (p *authzParser) Condition_() (localctx ICondition_Context) {
	localctx = NewCondition_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, authzParserRULE_condition_)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case authzParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(authzParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Match(authzParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(authzParserTIME_RANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case authzParserT__17:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Match(authzParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.Match(authzParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(authzParserLOCATION)
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
