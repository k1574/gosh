package token

//import "fmt"

type Type uint8

type Token struct {
	T Type
	V string
	L int
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
	Hashtag
	Escape
	If
)

func New(t Type, v string, l int) Token {
	return Token{
		T: t,
		V: v,
		L: l,}
}

func IsAnyOf(in Type, of []Type) bool {
	for _, v := range of {
		if v == in {
			return true
		}
	}

	return false
}

func RemoveAllOccurencesOf(t Type, from []Token) []Token {
	ret := []Token{}

	for _, v := range from {
		if v.T != t {
			ret = append(ret, v)
		}
	}

	return ret
}

