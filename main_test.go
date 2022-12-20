package main

import (
	"strings"
	"testing"

	"github.com/spatialcurrent/go-stringify/pkg/stringify"
)

func TestAdditions(t *testing.T) {
	input := `(module
				(func (export "addNumbers") (param i32 i32) (result i32)
					local.get 0
					local.get 1
					i32.add)
				)`
	_output := []interface{}{
		0, 97, 115, 109, 1, 0, 0, 0, 1, 7, 1, 96, 2, 127, 127, 1, 127, 3, 2, 1, 0, 7, 16, 1, 12, 34, 97, 100, 100, 78, 117, 109, 98, 101, 114, 115, 34, 0, 0, 10, 9, 1, 7, 0, 32, 0, 32, 1, 106, 11,
	}

	output := stringify.InterfaceSliceToStringSlice(_output)
	expectedOutput := strings.Join(output, " ")

	wasm := compile(input)

	if len(wasm) != len(expectedOutput) {
		t.Errorf("Expected %d, got %d", len(expectedOutput), len(wasm))
		return
	}

	if wasm != expectedOutput {
		t.Errorf("Expected %v, got %v", expectedOutput, wasm)
	}

}
func TestSubtractions(t *testing.T) {
	input := `(module
				(func (export "subNumbers") (param i32 i32) (result i32)
					local.get 0
					local.get 1
					i32.sub)
				)`
	_output := []interface{}{
		0, 97, 115, 109, 1, 0, 0, 0, 1, 7, 1, 96, 2, 127, 127, 1, 127, 3, 2, 1, 0, 7, 16, 1, 12, 34, 115, 117, 98, 78, 117, 109, 98, 101, 114, 115, 34, 0, 0, 10, 9, 1, 7, 0, 32, 0, 32, 1, 107, 11,
	}

	output := stringify.InterfaceSliceToStringSlice(_output)
	expectedOutput := strings.Join(output, " ")

	wasm := compile(input)

	if len(wasm) != len(expectedOutput) {
		t.Errorf("Expected %d, got %d", len(expectedOutput), len(wasm))
		return
	}

	if wasm != expectedOutput {
		t.Errorf("Expected %v, got %v", expectedOutput, wasm)
	}

}
func TestMultiplications(t *testing.T) {
	input := `(module
				(func (export "mulNumbers") (param i32 i32) (result i32)
					local.get 0
					local.get 1
					i32.mul)
				)`
	_output := []interface{}{
		0, 97, 115, 109, 1, 0, 0, 0, 1, 7, 1, 96, 2, 127, 127, 1, 127, 3, 2, 1, 0, 7, 16, 1, 12, 34, 109, 117, 108, 78, 117, 109, 98, 101, 114, 115, 34, 0, 0, 10, 9, 1, 7, 0, 32, 0, 32, 1, 108, 11,
	}

	output := stringify.InterfaceSliceToStringSlice(_output)
	expectedOutput := strings.Join(output, " ")

	wasm := compile(input)

	if len(wasm) != len(expectedOutput) {
		t.Errorf("Expected %d, got %d", len(expectedOutput), len(wasm))
		return
	}

	if wasm != expectedOutput {
		t.Errorf("Expected %v, got %v", expectedOutput, wasm)
	}

}
func TestDivisions(t *testing.T) {
	input := `(module
				(func (export "divNumbers") (param i32 i32) (result i32)
					local.get 0
					local.get 1
					i32.div)
				)`
	_output := []interface{}{
		0, 97, 115, 109, 1, 0, 0, 0, 1, 7, 1, 96, 2, 127, 127, 1, 127, 3, 2, 1, 0, 7, 16, 1, 12, 34, 100, 105, 118, 78, 117, 109, 98, 101, 114, 115, 34, 0, 0, 10, 9, 1, 7, 0, 32, 0, 32, 1, 109, 11,
	}

	output := stringify.InterfaceSliceToStringSlice(_output)
	expectedOutput := strings.Join(output, " ")

	wasm := compile(input)

	if len(wasm) != len(expectedOutput) {
		t.Errorf("Expected %d, got %d", len(expectedOutput), len(wasm))
		return
	}

	if wasm != expectedOutput {
		t.Errorf("Expected %v, got %v", expectedOutput, wasm)
	}

}
func TestConst(t *testing.T) {
	input := `(module
				(func (export "operationWithInternalVariable") (param i32 i32) (result i32)
					local.get 0
					i32.const 10
					i32.add)
				)
			`
	_output := []interface{}{
		0, 97, 115, 109, 1, 0, 0, 0, 1, 7, 1, 96, 2, 127, 127, 1, 127, 3, 2, 1, 0, 7, 35, 1, 31, 34, 111, 112, 101, 114, 97, 116, 105, 111, 110, 87, 105, 116, 104, 73, 110, 116, 101, 114, 110, 97, 108, 86, 97, 114, 105, 97, 98, 108, 101, 34, 0, 0, 10, 9, 1, 7, 0, 32, 0, 65, 10, 106, 11,
	}

	output := stringify.InterfaceSliceToStringSlice(_output)
	expectedOutput := strings.Join(output, " ")

	wasm := compile(input)

	if len(wasm) != len(expectedOutput) {
		t.Errorf("Expected %d, got %d", len(expectedOutput), len(wasm))
		return
	}

	if wasm != expectedOutput {
		t.Errorf("Expected %v, got %v", expectedOutput, wasm)
	}

}
