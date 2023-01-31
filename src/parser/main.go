package parser

import (
	"github.com/surdeus/gosh/src/token"
)

type Command []token.Token
type Type uint8

type Tree struct {
	V token.Token
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
}

