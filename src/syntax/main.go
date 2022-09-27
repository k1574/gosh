package syntax



const (
	Variable = '$'
	OpeningBrace = '{'
	ClosingBrace = '}'
	Quote = '\''
	Backquote = '`'
	Concat = '^' 
	Escape = '\\'
	Semicolon = ';'
	Ampersand = '&'
	Pipe = '|'
)

var (
	WordDels = []byte{' ', '\t',}
	SpecialChars = []byte{
		OpeningBrace,
		ClosingBrace,
		Semicolon,
		Ampersand,
		Pipe,
		Escape,
		Backquote,
		Concat}
)

func EqAnyOf[t byte | string](v t, a []t) bool {
	for _, c := range a {
		if v == c {
			return true
		}
	}

	return false
}

func TrimLeft[t byte](a []t, fn func(v t) bool) ([]t, []t) {
	var (
		i int
	)

	for i = 0 ;  i<len(a) && fn(a[i]); i++ {}

	return a[:i], a[i:]
}

func IsSpace(c byte) bool {
	return EqAnyOf[byte](c, WordDels)
}

func IsSpecial(c byte) bool {
	return EqAnyOf[byte](c, []byte{Concat,
		OpeningBrace,
		ClosingBrace,
		Escape,
		Semicolon})
}

func TrimLeftSpaces(s string) (string, string) {
	ret1, ret2 := TrimLeft[byte]([]byte(s), IsSpace)
	return string(ret1), string(ret2)
}

func TrimLeftWord(s string) (string, string) {
	ret1, ret2 := TrimLeft[byte]([]byte(s), func(b byte) bool {
		return !IsSpace(b) && !IsSpecial(b)
	})
	return string(ret1), string(ret2)
}

