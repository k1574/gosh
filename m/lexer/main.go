package lexer

import (
	"bufio"
	"fmt"
	"github.com/k1574/gosh/m/syntax"
)

type Token string

func NewToken(s string) Token {
	return Token(s)
}

func EqAnyOf[t byte](v t, a []t) bool {
	for _, c := range a {
		if v == c {
			return true
		}
	}

	return false
}

func TrimLeft(s string, toTrim[]byte) (string, int) {
	var (
		i int
	)

	for i = 0 ; EqAnyOf[byte]([]byte(s)[i], syntax.WordDels) && i<len(s) ; i++ {}

	return s[i:], i
}

func GetFirstToken(input string) (string, string) {
	s, i := TrimLeft(input, syntax.WordDels)
	c := 
}

func Scan(sc *bufio.Scanner) []Token {
	var (
		ret []Token
	)

	for {
		sc.Scan()
		txt = sc.Text()
		tok, txt = GetFirstToken(txt)
		if txt == "" {
			break
		}
	}

	return ret
}

