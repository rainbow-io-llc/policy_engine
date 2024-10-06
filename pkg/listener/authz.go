package listener

import (
	"fmt"

	authz "github.com/policy_engine/gen/antlr/authz"
)

type AuthZListener struct {
	*authz.BaseauthzListener
}

func NewAuthZListener() *AuthZListener {
	return &AuthZListener{}
}

func (al *AuthZListener) EnterRule(ctx *authz.RuleContext) {
	subject := ctx.Subject().GetText()
	action := ctx.Action_().GetText()
	object := ctx.Object().GetText()
	condition := ctx.Condition().GetText()

	fmt.Printf("Rule parsed: subject=%s, action=%s, object=%s, condition=%s\n", subject, action, object, condition)

	if ctx.Condition() != nil {
		time_range := ""
		location := ""
		for idx := range ctx.Condition().AllCondition_() {
			if ctx.Condition().Condition_(idx).TIME_RANGE() != nil {
				time_range = ctx.Condition().Condition_(idx).TIME_RANGE().GetText()
			}
			if ctx.Condition().Condition_(idx).LOCATION() != nil {
				location = ctx.Condition().Condition_(idx).LOCATION().GetText()
			}
		}
		fmt.Printf("TIME_RANGE: %s, LOCATION: %s\n", time_range, location)
	}
}