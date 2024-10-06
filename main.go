package main

import (
	"github.com/antlr4-go/antlr/v4"

	authz "github.com/policy_engine/gen/antlr/authz"
	"github.com/policy_engine/pkg/listener"
)

func main() {
	input := antlr.NewInputStream(`ALLOW service Exec collect TO data;`)

	lexer := authz.NewauthzLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := authz.NewauthzParser(stream)
	tree := parser.Policy()
	antlr.ParseTreeWalkerDefault.Walk(listener.NewAuthZListener(), tree)
}
