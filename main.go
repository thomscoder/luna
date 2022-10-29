package main

import (
	"fmt"
	"luna/compiler"
)

func main() {
	// Tokens
	tokens := compiler.Tokenize("log 34")
	fmt.Println("Tokens:", tokens)
	// Ast
	ast := compiler.Parser(tokens)
	fmt.Println("Ast:", ast)
}
