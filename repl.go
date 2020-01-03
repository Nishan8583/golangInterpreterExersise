package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("_______The Monkey Interpreter______")
	scanner := bufio.Scanner(os.Stdin)
	for {
		fmt.Printf(">>> ")
		line := scanner.Text()

	}
}
