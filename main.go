package main

import (
	"fmt"
	"hali/handlers"
	"hali/modules"
)

func main() {
	wasm := modules.Emitter()
	handlers.CreateFile(wasm)
	fmt.Println(wasm)
}
