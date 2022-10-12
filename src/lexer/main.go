package lexer

/* Implements basic separation of tokens in input. */

import (
	"errors"
	"github.com/surdeus/gosh/src/syntax"
	"github.com/surdeus/gosh/src/token"
	//"fmt"
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
	DeepLvl int
}

const (
	Free Status = iota
	InQuotedWord
	InBlock
	InComment
)

var (
	EOS = errors.New("end of string")
	NotFinishedQuotedWord = errors.New("not finished quoted word")
	ClosingBraceWithoutOpening = errors.New("closing brace without opening")
)

func New() *Lexer {
	var lexer Lexer
	lexer.Status = Free
	lexer.Handlers = Handlers {
		syntax.OpeningBrace : lexer.OpeningBrace,
		syntax.ClosingBrace : lexer.ClosingBrace,
		syntax.Quote : lexer.QuotedWord,
		syntax.Backquote : lexer.Backquote,
		syntax.Concat : lexer.Concat,
		syntax.Semicolon : lexer.Semicolon,
		syntax.Pipe : lexer.Pipe,
		syntax.Ampersand : lexer.Ampersand,
		syntax.Hashtag : lexer.Hashtag,
	}
   
	lexer.Tokens = []token.Token{}

	lexer.Storage = ""
	lexer.Line = 1
	lexer.DeepLvl = 0

	return &lexer
}

func (l *Lexer)QuotedWord(s string) (token.Token, string, error){
	var (
		i int
	)

	for i = 1 ; i < len(s) ; i++ {
		if s[i] == syntax.Quote {
			if i == (len(s)) - 1 {
				/* Last char in input is a Quote .*/
				return token.New(token.QuotedWord, s[1:i-1], l.Line), "", nil
			} else if s[i+1] != syntax.Quote {
				/* Found not escaped Quote. */
				return token.New(token.QuotedWord, s[1:i], l.Line), s[i+1:], nil
			}
		}
	}

	return token.New(token.QuotedWord, s[1:], l.Line), "", NotFinishedQuotedWord
}

func (l *Lexer)OpeningBrace(s string) (token.Token, string, error) {
	return token.New(token.OpeningBrace, s[0:1], l.Line), s[1:], nil
}

func (l *Lexer)ClosingBrace(s string) (token.Token, string, error) {
	return token.New(token.ClosingBrace, s[0:1], l.Line), s[1:], nil
}

func (l *Lexer)Backquote(s string) (token.Token, string, error) {
	return token.New(token.Backquote, s[0:1], l.Line), s[1:], nil
}

func (l *Lexer)Concat(s string) (token.Token, string, error) {
	return token.New(token.Concat, s[0:1], l.Line), s[1:], nil
}

func (l *Lexer)SimpleWord(s string) (token.Token, string, error){
	left, right := syntax.TrimLeftWord(s)

	return token.New(token.SimpleWord, left, l.Line), right, nil
}
func (l *Lexer)Semicolon(s string) (token.Token, string, error){
	return token.New(token.Semicolon, s[0:1], l.Line), s[1:], nil
}

func (l *Lexer)Ampersand(s string) (token.Token, string, error) {
	if len(s) > 1 {
		if s[1] == s[0] {
			return token.New(token.And, s[:2], l.Line), s[2:], nil
		} else if s[1] == syntax.Pipe {
			return token.New(token.If, s[:2], l.Line), s[2:], nil
		}
	} 

	return token.New(token.Background, s[:1], l.Line), s[1:], nil
}

func (l *Lexer)Pipe(s string) (token.Token, string, error) {
	if len(s) > 1 {
		if s[1] == s[0] {
			return token.New(token.Pipe, s[:2], l.Line), s[2:], nil
		} 
	}

	return token.New(token.Or, s[:1], l.Line), s[1:], nil
}

func (l *Lexer)Hashtag(s string) (token.Token, string, error) {
	return token.New(token.Hashtag, s[1:], l.Line), "", nil
}

func (l *Lexer)GetNextToken(input string) (token.Token, string, error) {
	_, s := syntax.TrimLeftSpaces(input)
	if len(s) == 0 {
		return token.New(token.Empty, "", l.Line), "", EOS
	}

	if v, notASimpleWord := l.Handlers[s[0]] ; notASimpleWord {
		return v(s)
	} else {
		return l.SimpleWord(s)
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

func (l *Lexer)Scan(txt string) (bool, error) {
	var (
		tok token.Token
		err error
	)

	if l.Status == InQuotedWord {
		caught, left, right := CatchFinishingQuote(txt)
		l.Storage += left
		if !caught {
			l.Storage += "\n"
			l.Line++
			return false, nil
		}
		l.Status = Free
		l.Tokens = append(l.Tokens, token.New(token.QuotedWord, l.Storage, l.Line))
		l.Storage = ""
		txt = right
	}


	for {
		tok, txt, err = l.GetNextToken(txt)
		if err == EOS {
			break
		} else if err == NotFinishedQuotedWord {
			l.Status = InQuotedWord
			l.Storage = tok.V + "\n"
			return false, nil
		} else if tok.T == token.OpeningBrace {
			l.DeepLvl++
		} else if tok.T == token.ClosingBrace {
			l.DeepLvl--
			if l.DeepLvl < 0 {
				return false, ClosingBraceWithoutOpening
			}
		} else if err != nil {
			return false, err
		}
		l.Tokens = append(l.Tokens, tok)
	}

	t := l.Tokens[len(l.Tokens)-1].T
	if !token.IsAnyOf(t, []token.Type{token.OpeningBrace,
			token.Semicolon,
			token.Escape,
			token.Background,
			token.Pipe,
			token.If, } ) {
		l.Tokens = append(l.Tokens, token.New(token.Semicolon, string(syntax.Semicolon), l.Line))
	}

	if t == token.Escape {
		l.Tokens = l.Tokens[:len(l.Tokens)-2]
	}

	l.Line++
	return l.Status == Free && l.DeepLvl == 0, nil
}

