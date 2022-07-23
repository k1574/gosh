package syntax

var (
	WordDels = []byte{'\t', ' '}
	OpDels = []rune{'\r', '\n', ';'}
	Quote = '\''
	CmdGroupOpen = '{'
	CmdGroupClose = '}'
	CmdOutputPref = '$'
	Pipe = ""
)
