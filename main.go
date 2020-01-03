package main

import (
	"log"

	"./lexer"
	"./token"
)

// TestNextToken will test if lexing works
func main() {
	log.Println("I am here")
	// the input form the command line
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x,y) {
		x+y;
	}
	== != 
	let result = add(five,ten)`
	log.Println("About to create a lexer object")
	l := lexer.New(input)
	log.Println("Finished creating a lexer object")

	for {
		tok := l.NextToken()
		if tok.Type == token.ILLEGAL {
			log.Println("EOF")
			break
		}
		log.Printf("The literal was %s and the token was %s", tok.Literal, tok.Type)

	}

}
