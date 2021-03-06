package main

import (
	"bufio"
	"fmt"
	"os"

	"./evaluator"
	"./lexer"
	"./object"
	"./parser"
)

func main() {
	banner := `                                                            
	/////(/(###########//////####////                
	////(###############(/////////////               
   ##/############////////////,,,,,,,/##             
  *##############//(/////////,,,,,,,,,,,/(######     
 *###############(///////////,,,,,,,,,,,,,,(####     
###############/////////////,,,,,,,,,,,.,,,,####     
###############(///////////,,,,,,,,,,,,....,*###     
# #####&&&&&&&&&##//////////%%%%&&%%////,...,/###     
  ##&&&&&#############/*/#%%((%%%&&&&&&&&&**(  %     
%&&              %&&&&%%&&#               *&#        
(&                 &&&#&&                  %*,/*     
###                &#((*.&                &.*.(*     
####&             &&##(**..%.             *..*.(*     
/#######%&&&&&%##&&###(**.../**/**//**,......*,/(     
//###################((**....,///****........*,((     
//######(((((###((((((******.*////**........*        
 #######((((####&&&##((((***********........*        
 /######(######&&&&##%%%#((*********.......*         
 //###################(((/**********....,,.(         
   ##########%&&&&%####((((****(##((******(%         
     ##############((((((((**************(#          
     ###################((((**********((             
       &%#############(((***...****((.               
       &&#%######(((((/******(#((.                
         ########(((######***(                
`
	fmt.Printf("___THE DADA Programming Interpreter___\n%s\nNOTE: This is not completely my own creation but will rather contain modification to the the monkey programming language as taught in the book in https://interpreterbook.com/ website\n\n",
		banner)
	reader := bufio.NewReader(os.Stdin)

	// Environment object to store the environment values
	env := object.NewEnvironment()
	for {
		fmt.Printf(">>> ")
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Error while reading input ->", err)
			return
		}
		if string(line) == "exit" || string(line) == "exit()" || string(line) == "quit" || string(line) == "quit()" {
			break
		}

		lex := lexer.New(string(line))
		p := parser.New(lex)

		programs := p.ParseProgram()
		/*
			fmt.Print("\nThe statementents from you input are\n\n")
			for index, value := range programs.Statements {
				fmt.Printf("$stmt>> %d		%v\n", index, value)
			}
			fmt.Println()
		*/
		evaluator.Debug("Programs: ", programs)
		evaluated := evaluator.Eval(programs, env)
		if evaluated != nil {
			fmt.Println(evaluated.Inspect())
		}

	}
}
