package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"luna/compiler"
	// "luna/handlers"
	// "luna/modules"
)

func main() {
	// wasm := modules.Emitter()
	// handlers.CreateFile(wasm)
	// fmt.Println(wasm)
	content, err := ioutil.ReadFile("../test/main.wat")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	fmt.Println(compiler.Parser(compiler.Tokenize("log 34")))
}
