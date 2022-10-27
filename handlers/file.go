package handlers

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateFile(data []uint8) {
	err := os.WriteFile("./main.wasm", data, 0644)
	check(err)
}
