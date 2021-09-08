package parser

import (
	"fmt"
	"strconv"

	"github.com/salleaffaire/ynt/ast"
	"github.com/salleaffaire/ynt/lexer"
	"github.com/salleaffaire/ynt/token"
)

type Parser struct {
	l *lexer.Lexer

	Errors []string

	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) Error(message string) {
	p.Errors = append(p.Errors, message)
}

func (p *Parser) printParserErrors() {
	for _, e := range p.Errors {
		fmt.Println(e)
	}
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		Errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseDocument() *ast.Document {
	document := &ast.Document{}
	document.Values = []ast.Value{}

	for p.curToken.Type != token.EOF {
		object := p.parseValue()
		if object != nil {
			document.Values = append(document.Values, object)
		} else {
			p.printParserErrors()
			return nil
		}

		p.nextToken()
	}

	return document
}

func (p *Parser) parseValue() ast.Value {
	switch p.curToken.Type {
	case token.NUMBER:
		return p.parseIntegerValue()
	case token.STRING:
		return p.parseStringValue()
	case token.TRUE:
		return p.parseBooleanValue()
	case token.FALSE:
		return p.parseBooleanValue()
	case token.LBRACKET:
		return p.parseArrayValue()
	case token.LBRACE:
		return p.parseObjectValue()
	default:
		msg := fmt.Sprintf("Error: unexpected token %s", p.curToken.Literal)
		p.Errors = append(p.Errors, msg)
		return nil
	}
}

func (p *Parser) parseIntegerValue() ast.Value {

	lit := &ast.NumberValue{Token: p.curToken}

	value, err := strconv.ParseFloat(p.curToken.Literal, 0)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.curToken.Literal)
		p.Errors = append(p.Errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) parseStringValue() ast.Value {

	lit := &ast.StringValue{Token: p.curToken}

	lit.Value = p.curToken.Literal

	return lit
}

func (p *Parser) parseArrayValue() ast.Value {
	arrayValue := &ast.ArrayValue{Token: p.curToken, Values: []ast.Value{}}

	if p.peekTokenIs(token.RBRACKET) {
		p.nextToken()
		return arrayValue
	}

	p.nextToken()
	value := p.parseValue()
	if value != nil {
		arrayValue.Values = append(arrayValue.Values, value)
	} else {
		return nil
	}

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		value := p.parseValue()
		if value != nil {
			arrayValue.Values = append(arrayValue.Values, value)
		} else {
			return nil
		}
	}

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return arrayValue
}

func (p *Parser) parseObjectValue() ast.Value {
	objectValue := &ast.ObjectValue{Token: p.curToken, Attributes: []ast.Attribute{}}

	if p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		return objectValue
	}

	att := ast.Attribute{}
	// Skip the left brace, curToken is the Key (STRING)
	p.nextToken()
	att.Key = p.curToken.Literal
	// Skip the key, curToken is a colon
	p.nextToken()
	// Skip the colon
	p.nextToken()
	att.V = p.parseValue()
	objectValue.Attributes = append(objectValue.Attributes, att)

	for p.peekTokenIs(token.COMMA) {
		att := ast.Attribute{}
		// Skip the curToken, curToken is now the comma
		p.nextToken()
		// Skip the comma, curToken is the Key (STRING)
		p.nextToken()
		att.Key = p.curToken.Literal
		// Skip the key, curToken is a colon
		p.nextToken()
		// Skip the colon
		p.nextToken()
		att.V = p.parseValue()
		objectValue.Attributes = append(objectValue.Attributes, att)
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return objectValue
}

func (p *Parser) parseBooleanValue() ast.Value {
	expression := &ast.BooleanValue{
		Token: p.curToken,
		Value: p.curTokenIs(token.TRUE),
	}
	return expression
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		// p.peekError(t)
		return false
	}
}
