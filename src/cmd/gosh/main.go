package main

import (
	"fmt"
	"github.com/surdeus/gosh/src/lexer"
	"bufio"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	l := lexer.New()
	scan := func() bool {
		fmt.Printf("> ")
		return sc.Scan()
	}
	for scan() {
		if finished, err := l.Scan(sc.Text()) ; !finished {
			if err != nil {
				fmt.Printf("error: line %d: %s\n", l.Line, err)
				break
			}
			continue
		}
		fmt.Printf("Status: %v\n", l.Status)
		fmt.Println("Tokens:")
		for _, v := range l.Tokens {
			fmt.Printf("%d\t\"%s\"\n", v.T, v.V)
		}
	}

	fmt.Println("")

	if err := sc.Err() ; err != nil {
		fmt.Printf("Error: '%s'\n", err)
		os.Exit(1)
	}
}

