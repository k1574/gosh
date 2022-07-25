package lexer

/* Implements basic separation of tokens in input. */

import (
	"bufio"
	//"fmt"
	"errors"
	"github.com/k1574/gosh/m/syntax"
)

var (
	Tokens = map[byte] func(string) (Token, string, error) {
		syntax.OpeningBrace : OpeningBrace,
		syntax.ClosingBrace : ClosingBrace,
		syntax.Quote : QuotedWord,
		syntax.CmdOutput : CmdOutput,
	}
	NotFinishedQuotedWord = errors.New("Not finished quoted word")
)

type Token string

func NewToken(s string) Token {
	return Token(s)
}

func QuotedWord(s string) (Token, string, error){
	var (
		i int
	)

	for i = 1 ; i < len(s) - 1 ; i++ {
		if s[i] == syntax.Quote{
			if i == (len(s)) - 1 {
				/* Last char in input is a Quote .*/
				return NewToken(s), "", nil
			} else if s[i+1] != syntax.Quote {
				/* Found not escaped Quote. */
				//fmt.Println("im in")
				return NewToken(s[:i+1]), s[i+1:], nil
			}
		}
	}

	return NewToken(""), s, NotFinishedQuotedWord
}

func OpeningBrace(s string) (Token, string, error){
	return NewToken(string(syntax.OpeningBrace)), s[1:], nil
}

func ClosingBrace(s string) (Token, string, error){
	return NewToken(string(syntax.ClosingBrace)), s[1:], nil
}

func CmdOutput(s string) (Token, string, error){
	return NewToken(string(syntax.CmdOutput)), s[1:], nil
}

func SimpleWord(s string) (Token, string, error){
	return NewToken(string(syntax.CmdOutput)), s[1:], nil
}

func GetNextToken(input string) (Token, string, error) {
	s, _ := syntax.TrimLeftSpaces(input)
	if len(s) == 0 {
		return NewToken(""), "", nil
	}

	if v, notASimpleWord := Tokens[s[0]] ; notASimpleWord {
		return v(s)
	} else {
		return SimpleWord(s)
	}
}

func Scan(sc *bufio.Scanner) ([]Token, error) {
	var (
		ret []Token
		tok Token
		err error
		txt string
	)

	sc.Scan()
	txt = sc.Text()
	for {
		tok, txt, err = GetNextToken(txt)
		//fmt.Printf("'%s' '%s'\n", string(tok), txt)
		if err != nil {
			break
		}
		ret = append(ret, tok)
		if txt == "" {
			break
		}
	}

	return ret, nil
}

