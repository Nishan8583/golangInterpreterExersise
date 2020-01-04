package main

import (
	"bufio"
	"fmt"
	"os"

	"./token"

	"./lexer"
)

func main() {
	fmt.Println("_______The Monkey Interpreter______")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">>> ")
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Error while reading input ->", err)
			return
		}

		lex := lexer.New(string(line))
		for {
			tok := lex.NextToken()
			if tok.Type == token.EOF {
				fmt.Println("End of token reached")
				break
			}
			fmt.Println(tok)
		}
	}
}
