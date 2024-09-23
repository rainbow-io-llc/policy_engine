// Code generated from antlr/authz.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // authz

import "github.com/antlr4-go/antlr/v4"

// BaseauthzListener is a complete listener for a parse tree produced by authzParser.
type BaseauthzListener struct{}

var _ authzListener = &BaseauthzListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseauthzListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseauthzListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseauthzListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseauthzListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPolicy is called when production policy is entered.
func (s *BaseauthzListener) EnterPolicy(ctx *PolicyContext) {}

// ExitPolicy is called when production policy is exited.
func (s *BaseauthzListener) ExitPolicy(ctx *PolicyContext) {}

// EnterRule is called when production rule is entered.
func (s *BaseauthzListener) EnterRule(ctx *RuleContext) {}

// ExitRule is called when production rule is exited.
func (s *BaseauthzListener) ExitRule(ctx *RuleContext) {}

// EnterSubject is called when production subject is entered.
func (s *BaseauthzListener) EnterSubject(ctx *SubjectContext) {}

// ExitSubject is called when production subject is exited.
func (s *BaseauthzListener) ExitSubject(ctx *SubjectContext) {}

// EnterAction is called when production action is entered.
func (s *BaseauthzListener) EnterAction(ctx *ActionContext) {}

// ExitAction is called when production action is exited.
func (s *BaseauthzListener) ExitAction(ctx *ActionContext) {}

// EnterObject is called when production object is entered.
func (s *BaseauthzListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseauthzListener) ExitObject(ctx *ObjectContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseauthzListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseauthzListener) ExitCondition(ctx *ConditionContext) {}
