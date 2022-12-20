package main

import (
	"fmt"
	"luna/compiler"
	"strings"

	// "syscall/js"

	"github.com/spatialcurrent/go-stringify/pkg/stringify"
)

func main() {
	mode := "broser"

	if mode != "browser" {
		input := `(module
				(func (export "operationWithInternalVariable") (param i32 i32) (result i32)
					local.get 0
					i32.const 10
					i32.add)
				)
			`
		compile(input)
		return
	}

	// // Spin it in browser
	// wait := make(chan struct{}, 0)
	// js.Global().Set("startLuna", js.FuncOf(startLuna))
	// <-wait
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

// // TINYGO NOTE:  there is no export as we registered this function in global
// func startLuna(this js.Value, args []js.Value) interface{} {
// 	input := args[0].String()
// 	return js.ValueOf(map[string]interface{}{
// 		"module": compile(input),
// 	})
// }
