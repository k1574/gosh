package token

//import "fmt"

type Type uint8

type Token struct {
	T Type
	V string
}

const (
	Empty Type = iota
	OpeningBrace
	ClosingBrace
	Backquote
	QuotedWord
	SimpleWord
	Concat
	Semicolon
	Pipe
	Or
	And
	Background
	Escape
	If
)

func New(t Type, v string) Token {
	return Token{
		T: t,
		V: v}
}

func IsAnyOf(in Type, of []Type) bool {
	for _, v := range of {
		if v == in {
			return true
		}
	}

	return false
}

