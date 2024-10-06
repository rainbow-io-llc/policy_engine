// Code generated from authz.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // authz

import "github.com/antlr4-go/antlr/v4"

// authzListener is a complete listener for a parse tree produced by authzParser.
type authzListener interface {
	antlr.ParseTreeListener

	// EnterPolicy is called when entering the policy production.
	EnterPolicy(c *PolicyContext)

	// EnterRule is called when entering the rule production.
	EnterRule(c *RuleContext)

	// EnterSubject is called when entering the subject production.
	EnterSubject(c *SubjectContext)

	// EnterAction is called when entering the action production.
	EnterAction(c *ActionContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterCondition_ is called when entering the condition_ production.
	EnterCondition_(c *Condition_Context)

	// ExitPolicy is called when exiting the policy production.
	ExitPolicy(c *PolicyContext)

	// ExitRule is called when exiting the rule production.
	ExitRule(c *RuleContext)

	// ExitSubject is called when exiting the subject production.
	ExitSubject(c *SubjectContext)

	// ExitAction is called when exiting the action production.
	ExitAction(c *ActionContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitCondition_ is called when exiting the condition_ production.
	ExitCondition_(c *Condition_Context)
}
