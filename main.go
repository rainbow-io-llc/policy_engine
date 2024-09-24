package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"

	authz "github.com/policy_engine/gen/antlr/authz"
)

func main() {
	input := antlr.NewInputStream(`ALLOW service FROM collect TO data WHEN time_range = "09:00-17:00";`)

	lexer := authz.NewauthzLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	/* print token
	fmt.Println("\nLexer Tokens:")
    token := lexer.NextToken()
    for token.GetTokenType() != antlr.TokenEOF {
        fmt.Printf("Token: %s (Type: %d, Text: %s)\n", lexer.SymbolicNames[token.GetTokenType()], token.GetTokenType(), token.GetText())
        token = lexer.NextToken()
    }
	*/

	parser := authz.NewauthzParser(stream)
	tree := parser.Policy()

	fmt.Println("\nParser Tree:")
	fmt.Println(tree.ToStringTree(nil, parser))
}
