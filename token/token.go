package token

type Tokentype string


type Token struct {
	Type Tokentype
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 1343456
	// Operators
	ASSIGN = "="
	PLUS = "+"
	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	)

	var Keywords = map[string]Tokentype {
		"fn": FUNCTION,
		"let": LET,
	}

	func LookupIdent(ident string) Tokentype{
		if tok, ok := Keywords[ident]; ok {
			return tok
		}
		return IDENT
	}

	