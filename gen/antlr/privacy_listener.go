// Code generated from antlr/privacy.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // privacy

import "github.com/antlr4-go/antlr/v4"

// privacyListener is a complete listener for a parse tree produced by privacyParser.
type privacyListener interface {
	antlr.ParseTreeListener

	// EnterPolicy is called when entering the policy production.
	EnterPolicy(c *PolicyContext)

	// EnterRule is called when entering the rule production.
	EnterRule(c *RuleContext)

	// EnterAction is called when entering the action production.
	EnterAction(c *ActionContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// ExitPolicy is called when exiting the policy production.
	ExitPolicy(c *PolicyContext)

	// ExitRule is called when exiting the rule production.
	ExitRule(c *RuleContext)

	// ExitAction is called when exiting the action production.
	ExitAction(c *ActionContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)
}
