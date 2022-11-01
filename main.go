package main

import (
	"fmt"
	"luna/compiler"
	"strings"
	"syscall/js"

	"github.com/spatialcurrent/go-stringify/pkg/stringify"
)

func main() {
	mode := "browser"

	if mode != "browser" {
		input := `
			module
			function 
			param i32 i32 
			get 0
			get 1
			i32.add
			result i32
			export "addNumbers" 
		`
		compile(input)
		return
	}

	// Spin it in browser
	wait := make(chan struct{}, 0)
	js.Global().Set("startLuna", js.FuncOf(startLuna))
	<-wait
}

//export compile
func compile(input string) string {
	// c, err := ioutil.ReadFile()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// content := string(c)
	// fmt.Println(content)
	// Tokens
	tokens := compiler.Tokenize(input)
	fmt.Println("Tokens:", tokens)
	fmt.Println("----------------------------------------------------------------")
	// Ast
	ast := compiler.Parser(tokens)
	fmt.Println("Ast:", ast)

	// Emitters
	wasm := compiler.Compile(ast)
	fmt.Println("Wasm", wasm)

	// Str for Javascript
	str := stringify.InterfaceSliceToStringSlice(wasm)
	return strings.Join(str, " ")
}

// TINYGO NOTE:  there is no export as we registered this function in global
func startLuna(this js.Value, args []js.Value) interface{} {
	input := args[0].String()
	return js.ValueOf(map[string]interface{}{
		"module": compile(input),
	})
}
