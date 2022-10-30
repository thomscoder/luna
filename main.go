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
	// Encoding Unsigned
	fmt.Println("Unsigned Encoding:", compiler.EncodeUnsignedLEB128(34))
	// Encoding Signed
	fmt.Println("Signed Encoding:", compiler.EncodeSignedLEB128(82))
	// Buffer Encoding
	fmt.Println("Buffer Encoding:", compiler.Ieee754(98))

	// Emitters
	fmt.Println("Emitter", compiler.Emitter(ast))
}
