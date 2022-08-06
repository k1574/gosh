package token

type Type uint8

type Token struct {
	T Type
	V string
}

const (
	Empty Type = iota
	OpeningBrace
	ClosingBrace
	CmdOutput
	QuotedWord
	SimpleWord
	Concat
	Semicolon
	Pipe
	Or
	And
	Background
	Escape
)

func New(t Type, v string) Token {
	return Token{t, v}
}

