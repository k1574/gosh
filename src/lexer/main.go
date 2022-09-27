package lexer

/* Implements basic separation of tokens in input. */

import (
	"errors"
	"github.com/surdeus/gosh/src/syntax"
	"github.com/surdeus/gosh/src/token"
	"fmt"
)

var (
	Tokens = map[byte] func(string) (token.Token, string, error) {
		syntax.OpeningBrace : OpeningBrace,
		syntax.ClosingBrace : ClosingBrace,
		syntax.Quote : QuotedWord,
		syntax.Backquote : Backquote,
		syntax.Concat : Concat,
		syntax.Semicolon : Semicolon,
		syntax.Pipe : Pipe,
		syntax.Ampersand : Ampersand,
		syntax.Escape : Escape,
	}
	NotFinishedQuotedWord = errors.New("Not finished quoted word")
	EOS = errors.New("end of string")
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

	return token.New(token.Empty, ""), s, NotFinishedQuotedWord
}

func OpeningBrace(s string) (token.Token, string, error) {
	return token.New(token.OpeningBrace, s[0:1]), s[1:], nil
}

func ClosingBrace(s string) (token.Token, string, error) {
	return token.New(token.ClosingBrace, s[0:1]), s[1:], nil
}

func Backquote(s string) (token.Token, string, error) {
	return token.New(token.Backquote, s[0:1]), s[1:], nil
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
func Escape(s string) (token.Token, string, error){
	return token.New(token.Escape, s[0:1]), s[1:], nil
}

func Ampersand(s string) (token.Token, string, error) {
	if len(s) > 1 {
		if s[1] == s[0] {
			return token.New(token.And, s[:2]), s[2:], nil
		} else if s[1] == syntax.Pipe {
			fmt.Println("im in")
			return token.New(token.If, s[:2]), s[2:], nil
		}
	} 

	return token.New(token.Background, s[:1]), s[1:], nil
}

func Pipe(s string) (token.Token, string, error) {
	if len(s) > 1 {
		if s[1] == s[0] {
			return token.New(token.Pipe, s[:2]), s[2:], nil
		} 
	}

	return token.New(token.Or, s[:1]), s[1:], nil
}

func GetNextToken(input string) (token.Token, string, error) {
	_, s := syntax.TrimLeftSpaces(input)
	if len(s) == 0 {
		return token.New(token.Empty, ""), "", EOS
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
		if err == EOS {
			break
		} else if err != nil {
			return []token.Token{}, err
		}
		ret = append(ret, tok)
	}

	t := ret[len(ret)-1].T
	if !token.IsAnyOf(t, []token.Type{token.OpeningBrace,
			token.ClosingBrace,
			token.Escape} ) {
		ret = append(ret, token.New(token.Semicolon, string(syntax.Semicolon)))
	}

	if t == token.Escape {
		ret = ret[:len(ret)-2]
	}

	return ret, nil
}

