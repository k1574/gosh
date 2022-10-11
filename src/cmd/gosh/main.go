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
	for sc.Scan() {
		l.Scan(sc.Text())
		fmt.Printf("Status: %v\n", l.Status)
		fmt.Println("Tokens:")
		for _, v := range l.Tokens {
			fmt.Printf("%d\t\"%s\"\n", v.T, v.V)
		}
	}

	if err := sc.Err() ; err != nil {
		fmt.Printf("Error: '%s'\n", err)
		os.Exit(1)
	}
}

