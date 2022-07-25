package token

type Type int

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
)

func New(t Type, v string) Token {
	return Token{t, v}
}

