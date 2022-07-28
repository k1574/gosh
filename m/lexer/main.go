package lexer

/* Implements basic separation of tokens in input. */

import (
	//"bufio"
	//"fmt"
	"errors"
	"github.com/k1574/gosh/m/syntax"
	"github.com/k1574/gosh/m/token"
)

var (
	Tokens = map[byte] func(string) (token.Token, string, error) {
		syntax.OpeningBrace : OpeningBrace,
		syntax.ClosingBrace : ClosingBrace,
		syntax.Quote : QuotedWord,
		syntax.CmdOutput : CmdOutput,
		syntax.Concat : Concat,
		syntax.Semicolon : Semicolon,
	}
	NotFinishedQuotedWord = errors.New("Not finished quoted word")
)

func QuotedWord(s string) (token.Token, string, error){
	var (
		i int
	)

	for i = 1 ; i < len(s) - 1 ; i++ {
		if s[i] == syntax.Quote {
			if i == (len(s)) - 1 {
				/* Last char in input is a Quote .*/
				return token.New(token.QuotedWord, s[1:i-1]), "", nil
			} else if s[i+1] != syntax.Quote {
				/* Found not escaped Quote. */
				return token.New(token.QuotedWord, s[1:i]), s[i+1:], nil
			}
		}
	}

	return token.New(token.Error, ""), s, NotFinishedQuotedWord
}

func OpeningBrace(s string) (token.Token, string, error) {
	return token.New(syntax.OpeningBrace, s[0:1]), s[1:], nil
}

func ClosingBrace(s string) (token.Token, string, error) {
	return token.New(token.OpeningBrace, s[0:1]), s[1:], nil
}

func CmdOutput(s string) (token.Token, string, error) {
	return token.New(token.CmdOutput, s[0:1]), s[1:], nil
}

func Concat(s string) (token.Token, string, error) {
	return token.New(token.Concat, s[0:1]), s[1:], nil
}

func SimpleWord(s string) (token.Token, string, error){
	left, right := syntax.TrimLeftWord(s)

	return token.New(token.SimpleWord, left), right, nil
}
func Semicolon(s string) (token.Token, string, error){
	return token.New(token.Semicolon, s[0:1]), s[1:], nil
}

func GetNextToken(input string) (token.Token, string, error) {
	_, s := syntax.TrimLeftSpaces(input)
	if len(s) == 0 {
		return token.New(token.Empty, ""), "", nil
	}

	if v, notASimpleWord := Tokens[s[0]] ; notASimpleWord {
		return v(s)
	} else {
		return SimpleWord(s)
	}
}

func Scan(txt string) ([]token.Token, error) {
	var (
		ret []token.Token
		tok token.Token
		err error
	)

	for {
		tok, txt, err = GetNextToken(txt)
		if err != nil {
			return []token.Token{}, err
		}
		ret = append(ret, tok)
		if txt == "" {
			break
		}
	}

	ret = append(ret, token.New(token.Semicolon, string(syntax.Semicolon)))

	return ret, nil
}

