package lexer

import (
	"fmt"
	"testing"

	"github.com/salleaffaire/ynt/token"
)

func TestNextToken(t *testing.T) {
	input := `{
		"a" : {
			"foo" : {
				"bar" : [0,2,3,4],
				"foobar" : 1
				"jj" : false
			}
		}
	}
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "a"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.STRING, "bar"},
		{token.COLON, ":"},
		{token.LBRACKET, "["},
		{token.NUMBER, "0"},
		{token.COMMA, ","},
		{token.NUMBER, "2"},
		{token.COMMA, ","},
		{token.NUMBER, "3"},
		{token.COMMA, ","},
		{token.NUMBER, "4"},
		{token.RBRACKET, "]"},
		{token.COMMA, ","},
		{token.STRING, "foobar"},
		{token.COLON, ":"},
		{token.NUMBER, "1"},
		{token.STRING, "jj"},
		{token.COLON, ":"},
		{token.FALSE, "false"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		// fmt.Println("Test : ", tt.expectedType)
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - litearl wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}

		fmt.Println("Token Type   : ", tok.Type)
		fmt.Println("Token Literal: ", tok.Literal)
		fmt.Println()

	}
}
