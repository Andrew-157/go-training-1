package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a mathematical expression: ")
		if eof := !scanner.Scan(); eof { // handle ctrl+d
			fmt.Println()
			os.Exit(0)
		}
		expression := scanner.Text()
		if strings.ToLower(expression) == "exit" {
			os.Exit(0)
		}
		fmt.Printf("You entered: %s\n", expression)
	}
}
