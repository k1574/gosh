package parser

import (
	"github.com/surdeus/gosh/src/token"
)

type Command []token.Token
type Type uint8

type Tree struct {
	T Type
	// Value
	V string
	// Children
	C []Tree
}

const (
	Root Type = iota
	Block
	If
	Else
	Or
	And
	Command
	Backquote
	SimpleWord
	QuotedWord
)

func Parse(tk []token.Token) (Tree, error) {
	var (
		i int
		t Tree = Tree{}
	)

	for i = 0 ; i < len(tk) ; i++ {
	}

	return t
}

func FindByTypeFromEnd(tk []token.Token, f token.Type) int {
	var (
		i int
	)

	for i = len(tk) - 1 ; i>=0 ; i-- {
		if tk[i].T == f {
			break
		}
	}

	return i
}

