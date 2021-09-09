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

func TestNextTokenNumber(t *testing.T) {

	tests := []struct {
		input    string
		expected []struct {
			Type    token.TokenType
			Literal string
		}
	}{
		{"000", []struct {
			Type    token.TokenType
			Literal string
		}{
			{token.NUMBER, "0"},
			{token.NUMBER, "0"},
			{token.NUMBER, "0"},
			{token.EOF, ""},
		}},
		{"{1,true}", []struct {
			Type    token.TokenType
			Literal string
		}{
			{token.LBRACE, "{"},
			{token.NUMBER, "1"},
			{token.COMMA, ","},
			{token.TRUE, "true"},
			{token.RBRACE, "}"},
			{token.EOF, ""},
		}},
		{"0.48e-8", []struct {
			Type    token.TokenType
			Literal string
		}{
			{token.NUMBER, "0.48e-8"},
			{token.EOF, ""},
		}},
		{"-0.48E8", []struct {
			Type    token.TokenType
			Literal string
		}{
			{token.NUMBER, "-0.48E8"},
			{token.EOF, ""},
		}},
		{"-1000000000000000000000000000", []struct {
			Type    token.TokenType
			Literal string
		}{
			{token.NUMBER, "-1000000000000000000000000000"},
			{token.EOF, ""},
		}},
		{"-1000000000000000000000000000", []struct {
			Type    token.TokenType
			Literal string
		}{
			{token.NUMBER, "-1000000000000000000000000000"},
			{token.EOF, ""},
		}},
		{"\"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*)(\\\\\t\n\b\f\r0123456789\"", []struct {
			Type    token.TokenType
			Literal string
		}{
			{token.STRING, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*)(\\\\\t\n\b\f\r0123456789"},
			{token.EOF, ""},
		}},
	}

	for i, tt := range tests {
		l := New(tt.input)

		for j, e := range tt.expected {
			tok := l.NextToken()

			if tok.Type != e.Type {
				t.Fatalf("tests[%d] - tokentype [%d] wrong. expected=%q, got=%q",
					i, j, e.Type, tok.Type)
			}

			if tok.Literal != e.Literal {
				t.Fatalf("tests[%d] - literal [%d] wrong. expected=%q, got=%q",
					i, j, e.Literal, tok.Literal)
			}

			fmt.Println("Token Type   : ", tok.Type)
			fmt.Println("Token Literal: ", tok.Literal)
			fmt.Println()
		}

	}
}
