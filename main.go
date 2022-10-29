package main

import (
	"fmt"
	"luna/compiler"
	// "luna/handlers"
	// "luna/modules"
)

func main() {
	// wasm := modules.Emitter()
	// handlers.CreateFile(wasm)
	// fmt.Println(wasm)

	fmt.Println(compiler.Tokenize("log 34"))
}
