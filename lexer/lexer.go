package lexer

import (
	"fmt"

	"github.com/salleaffaire/ynt/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte

	// Keep token dependent state
	state bool

	// Keep track of error message
	Errors []string

	lineNumber int

	// Tokens
	Tokens         []token.Token
	tokenIndex     int
	numberOfTokens int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.tokenIndex = 0
	l.lineNumber = 1

	l.readChar()

	tok := l.nextToken()
	if tok.Type == token.ILLEGAL {
		l.printLexerErrors()
		return nil
	}
	l.Tokens = append(l.Tokens, tok)

	for tok.Type != token.EOF {
		tok = l.nextToken()
		if tok.Type == token.ILLEGAL {
			l.printLexerErrors()
			return nil
		}
		l.Tokens = append(l.Tokens, tok)
	}

	l.numberOfTokens = len(l.Tokens)

	return l
}

func (l *Lexer) Error(message string) {
	l.Errors = append(l.Errors, message)
}

func (l *Lexer) printLexerErrors() {
	for _, e := range l.Errors {
		fmt.Println(e)
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
	// fmt.Println("l.ch: ", string(l.ch))
}

func (l *Lexer) NextToken() token.Token {
	tok := token.Token{}

	if l.tokenIndex >= l.numberOfTokens {
		tok.Type = token.EOF
	} else {
		tok = l.Tokens[l.tokenIndex]
	}
	l.tokenIndex += 1

	return tok
}

func (l *Lexer) nextToken() token.Token {
	var tok token.Token

	// fmt.Println("l.ch before white space: ", string(l.ch))
	l.skipWhitespace()
	// fmt.Println("l.ch after white space: ", string(l.ch))

	switch l.ch {
	case ':':
		tok = newToken(token.COLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)

	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
		if len(l.Errors) != 0 {
			tok.Type = token.ILLEGAL
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.NUMBER
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readString() string {
	position := l.position + 1

	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
		if l.ch == 0 {
			mes := fmt.Sprintf("Error: unexpected end of file in string %s - line %d position %d",
				l.input[position:l.position], l.lineNumber, l.position)
			l.Error(mes)
			break
		}
		if l.ch == '\\' {
			l.readChar()
			if (l.ch == '"') ||
				(l.ch == '/') ||
				(l.ch == '\\') ||
				(l.ch == 'b') ||
				(l.ch == 'f') ||
				(l.ch == 'n') ||
				(l.ch == 'r') ||
				(l.ch == 't') {
			} else {
				mes := fmt.Sprintf("Error: unexpected caracter %c in string %s - line %d position %d",
					l.ch, l.input[position:l.position], l.lineNumber, l.position)
				l.Error(mes)
				break
			}
		}
	}
	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	l.state = false
	for isDigit(l.ch) || l.ch == '.' {
		// fmt.Println("Digit: ", string(l.ch))
		l.readChar()
		if l.ch == '.' {
			l.state = true
		}
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' || l.ch == '\r' {
			l.lineNumber += 1
		}
		l.readChar()
	}
}
