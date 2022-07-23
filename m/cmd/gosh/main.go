package main

import (
	"fmt"
	"github.com/k1574/gosh/m/lexer"
	"bufio"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("'%s'\n", lexer.Scan(sc))
	}
}

