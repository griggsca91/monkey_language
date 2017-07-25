// token/token.go

package token

// Token - TODO
type Type string

// Token - TODO
type Token struct {
	Type    Type
	Literal string
}

const (
	Illegal = "Illegal"
	EOF     = "EOF"

	Ident = "Ident" // add, foobar, x, y, ....
	Int   = "Int"   // 1342

	Assign   = "="
	Plus     = "+"
	Minus    = "-"
	Bang    = "!"
	Asterisk = "*"
	Slash    = "/"

	Lt  = "<"
	Gt  = ">"
	Eq  = "=="
	Neq = "!="

	Comma     = ","
	Semicolon = ";"

	LParen = "("
	RParen = ")"
	LBrace = "{"
	RBrace = "}"

	Function = "Function"
	Let      = "Let"
	If       = "If"
	Else     = "Else"
	True     = "True"
	False    = "False"
	Return   = "Return"
)

var keywords = map[string]Type{
	"fn":     Function,
	"let":    Let,
	"if":     If,
	"else":   Else,
	"return": Return,
	"true":   True,
	"false":  False,
}

// LookupIdent - TODO
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Ident
}
