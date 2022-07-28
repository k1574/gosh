package token

type Type uint8

type Token struct {
	T Type
	V string
}

const (
	Error Type = iota
	Empty
	OpeningBrace
	ClosingBrace
	CmdOutput
	QuotedWord
	SimpleWord
	Concat
	Semicolon
)

func New(t Type, v string) Token {
	return Token{t, v}
}

