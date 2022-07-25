package syntax

var (
	WordDels = []byte{' ', '\t',}
	SpecialChars = []byte{'$', '{', '}', '`',}
)

const (
	Variable = '$'
	OpeningBrace = '{'
	ClosingBrace = '}'
	Quote = '\''
	CmdOutput = '`'
	Concat = '^' 
	Escape = '\\'
)

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

	for i = 0 ; EqAnyOf[byte]([]byte(s)[i], toTrim) && i<len(s) ; i++ {}

	return s[i:], i
}

func TrimLeftSpaces(s string) (string, int) {
	return TrimLeft(s, WordDels)
}

