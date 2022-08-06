package parser

import (
	"github.com/surdeus/gosh/src/token"
)

type Command []token.Token
type Type uint8

type Tree struct {
	T Type
	V string
	Ops []Tree
}

const (
	Root Type = iota

)

var (
	Operations = map[token.Type] Operation {
		token.And : Operation{ 2 },
		token.Or : Operation{ 1 },
		token.Pipe : Operation{0},
		token.Background : Operation{-1}
	}
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

