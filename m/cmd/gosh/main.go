package main

import (
	"fmt"
	"github.com/k1574/gosh/m/lexer"
	"bufio"
	"os"
	"io"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		out, err := lexer.Scan(sc.Text())
		if err != nil {
			break
		}
		fmt.Println("Tokens:")
		for _, v := range out {
			fmt.Printf("%v\n", v)
		}
	}

	if err := sc.Err() ; err != nil {
		fmt.Printf("Error: '%s'\n", err)
		os.Exit(1)
	}
}

