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
		out, err := lexer.Scan(sc)
		if err != nil {
			fmt.Printf("Error: '%s'\n", err)
		}
		fmt.Println("Tokens:")
		for _, v := range out {
			fmt.Printf("'%s'\n", v)
		}
	}
}

