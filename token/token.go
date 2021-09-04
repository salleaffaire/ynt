package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"
	NUMBER = "NUMBER"
	STRING = "STRING"

	// Delimiters
	COMMA = ","
	COLON = ":"

	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	TRUE  = "TRUE"
	FALSE = "FALSE"
)

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(s string) TokenType {
	if tt, ok := keywords[s]; ok {
		return tt
	}
	return IDENT
}
