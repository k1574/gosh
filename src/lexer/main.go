package lexer

/* Implements basic separation of tokens in input. */

import (
	"errors"
	"github.com/surdeus/gosh/src/syntax"
	"github.com/surdeus/gosh/src/token"
	"fmt"
)

type Status int8
type Handler = func(string) (token.Token, string, error)
type Handlers map[byte] Handler

type Lexer struct {
	Status Status
	Handlers Handlers
	Tokens []token.Token
	Storage string
	Line int
}

const (
	Free Status = iota
	InQuotedWord
	InBlock
	InComment
)

var (
	EOS = errors.New("end of string")
	NotFinishedQuotedWord = errors.New( "not finished quoted word")
)

func New() *Lexer {
	var lexer Lexer
	lexer.Status = Free
	lexer.Handlers = Handlers {
		syntax.OpeningBrace : OpeningBrace,
		syntax.ClosingBrace : ClosingBrace,
		syntax.Quote : QuotedWord,
		syntax.Backquote : Backquote,
		syntax.Concat : Concat,
		syntax.Semicolon : Semicolon,
		syntax.Pipe : Pipe,
		syntax.Ampersand : Ampersand,
		syntax.Escape : Escape,
		syntax.Hashtag : Hashtag,
	}

	lexer.Tokens = []token.Token{}

	lexer.Storage = ""
	lexer.Line = 1

	return &lexer
}

func QuotedWord(s string) (token.Token, string, error){
	var (
		i int
	)

	for i = 1 ; i < len(s) ; i++ {
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

	return token.New(token.QuotedWord, s[1:]), "", NotFinishedQuotedWord
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

func Hashtag(s string) (token.Token, string, error) {
	return token.New(token.Hashtag, s[1:]), "", nil
}

func (l *Lexer)GetNextToken(input string) (token.Token, string, error) {
	_, s := syntax.TrimLeftSpaces(input)
	if len(s) == 0 {
		return token.New(token.Empty, ""), "", EOS
	}

	if v, notASimpleWord := l.Handlers[s[0]] ; notASimpleWord {
		return v(s)
	} else {
		return SimpleWord(s)
	}
}

func CatchFinishingQuote(s string) (bool, string, string) {
	for i, v := range s {
		if v == syntax.Quote {
			// End of string.
			if i == len(s)-1 || s[i+1] != syntax.Quote {
				return true, s[:i], s[i+1:]
			}
		}
	}
	return false, s, ""
}

func (l *Lexer)Scan(txt string) bool {
	var (
		tok token.Token
		err error
	)

	if l.Status == InQuotedWord {
		caught, left, right := CatchFinishingQuote(txt)
		l.Storage += left
		fmt.Println(caught)
		if !caught {
			l.Storage += "\n"
			return false
		}
		l.Status = Free
		fmt.Println(l.Status)
		l.Tokens = append(l.Tokens, token.New(token.QuotedWord, l.Storage))
		l.Storage = ""
		txt = right
	}

	fmt.Println(txt)

	for {
		tok, txt, err = l.GetNextToken(txt)
		if err == EOS {
			break
		} else if err == NotFinishedQuotedWord {
			fmt.Println("getting not finished")
			l.Status = InQuotedWord
			l.Storage = tok.V + "\n"
			return false
		}else if err != nil {
			return false
		}
		fmt.Println("adding")
		l.Tokens = append(l.Tokens, tok)
	}

	t := l.Tokens[len(l.Tokens)-1].T
	if !token.IsAnyOf(t, []token.Type{token.OpeningBrace,
			token.Semicolon,
			token.Escape} ) {
		l.Tokens = append(l.Tokens, token.New(token.Semicolon, string(syntax.Semicolon)))
	}

	if t == token.Escape {
		l.Tokens = l.Tokens[:len(l.Tokens)-2]
	}

	return l.Status == Free
}

