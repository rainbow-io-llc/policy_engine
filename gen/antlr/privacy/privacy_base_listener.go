// Code generated from privacy.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // privacy

import "github.com/antlr4-go/antlr/v4"

// BaseprivacyListener is a complete listener for a parse tree produced by privacyParser.
type BaseprivacyListener struct{}

var _ privacyListener = &BaseprivacyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseprivacyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseprivacyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseprivacyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseprivacyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPolicy is called when production policy is entered.
func (s *BaseprivacyListener) EnterPolicy(ctx *PolicyContext) {}

// ExitPolicy is called when production policy is exited.
func (s *BaseprivacyListener) ExitPolicy(ctx *PolicyContext) {}

// EnterRule is called when production rule is entered.
func (s *BaseprivacyListener) EnterRule(ctx *RuleContext) {}

// ExitRule is called when production rule is exited.
func (s *BaseprivacyListener) ExitRule(ctx *RuleContext) {}

// EnterAction is called when production action is entered.
func (s *BaseprivacyListener) EnterAction(ctx *ActionContext) {}

// ExitAction is called when production action is exited.
func (s *BaseprivacyListener) ExitAction(ctx *ActionContext) {}

// EnterTarget is called when production target is entered.
func (s *BaseprivacyListener) EnterTarget(ctx *TargetContext) {}

// ExitTarget is called when production target is exited.
func (s *BaseprivacyListener) ExitTarget(ctx *TargetContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseprivacyListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseprivacyListener) ExitCondition(ctx *ConditionContext) {}

// EnterCondition_ is called when production condition_ is entered.
func (s *BaseprivacyListener) EnterCondition_(ctx *Condition_Context) {}

// ExitCondition_ is called when production condition_ is exited.
func (s *BaseprivacyListener) ExitCondition_(ctx *Condition_Context) {}
