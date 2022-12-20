package compiler

import (
	"luna/types"
	"testing"
)

// Test additions
func TestAddition(t *testing.T) {
	input := `(module
				(func (export "addNumbers") (param i32 i32) (result i32)
					local.get 0
					local.get 1
					i32.add)
				)`
	expectedOutput := []types.Token{
		{Type: "token", Value: "module", Index: 5},
		{Type: "token", Value: "func", Index: 17},
		{Type: "token", Value: "export", Index: 23},
		{Type: "literal", Value: "\"addNumbers\"", Index: 30},
		{Type: "token", Value: "param", Index: 45},
		{Type: "typeNum", Value: "i32", Index: 51},
		{Type: "typeNum", Value: "i32", Index: 55},
		{Type: "token", Value: "result", Index: 61},
		{Type: "typeNum", Value: "i32", Index: 68},
		{Type: "instruction", Value: "local.get", Index: 78},
		{Type: "number", Value: "0", Index: 88},
		{Type: "instruction", Value: "local.get", Index: 95},
		{Type: "number", Value: "1", Index: 105},
		{Type: "instruction", Value: "i32.add", Index: 112},
	}

	tokens := Tokenize(input)

	if len(tokens) != len(expectedOutput) {
		t.Errorf("Expected %d, got %d", len(expectedOutput), len(tokens))
		return
	}

	for i, token := range tokens {
		if token.Value != expectedOutput[i].Value {
			t.Errorf("Expected token value %s, got %s", expectedOutput[i].Value, token.Value)
		}
	}
}

// Test Subtractions
func TestSubtractions(t *testing.T) {
	input := `(module
				(func (export "subNumbers") (param i32 i32) (result i32)
					local.get 0
					local.get 1
					i32.sub)
				)`
	expectedOutput := []types.Token{
		{Type: "token", Value: "module", Index: 5},
		{Type: "token", Value: "func", Index: 17},
		{Type: "token", Value: "export", Index: 23},
		{Type: "literal", Value: "\"subNumbers\"", Index: 30},
		{Type: "token", Value: "param", Index: 45},
		{Type: "typeNum", Value: "i32", Index: 51},
		{Type: "typeNum", Value: "i32", Index: 55},
		{Type: "token", Value: "result", Index: 61},
		{Type: "typeNum", Value: "i32", Index: 68},
		{Type: "instruction", Value: "local.get", Index: 78},
		{Type: "number", Value: "0", Index: 88},
		{Type: "instruction", Value: "local.get", Index: 95},
		{Type: "number", Value: "1", Index: 105},
		{Type: "instruction", Value: "i32.sub", Index: 112},
	}

	tokens := Tokenize(input)

	if len(tokens) != len(expectedOutput) {
		t.Errorf("Expected %d, got %d", len(expectedOutput), len(tokens))
		return
	}

	for i, token := range tokens {
		if token.Value != expectedOutput[i].Value {
			t.Errorf("Expected token value %s, got %s", expectedOutput[i].Value, token.Value)
		}
	}
}

// Test Multiplications
func TestMultiplications(t *testing.T) {
	input := `(module
				(func (export "mulNumbers") (param i32 i32) (result i32)
					local.get 0
					local.get 1
					i32.mul)
				)`
	expectedOutput := []types.Token{
		{Type: "token", Value: "module", Index: 5},
		{Type: "token", Value: "func", Index: 17},
		{Type: "token", Value: "export", Index: 23},
		{Type: "literal", Value: "\"mulNumbers\"", Index: 30},
		{Type: "token", Value: "param", Index: 45},
		{Type: "typeNum", Value: "i32", Index: 51},
		{Type: "typeNum", Value: "i32", Index: 55},
		{Type: "token", Value: "result", Index: 61},
		{Type: "typeNum", Value: "i32", Index: 68},
		{Type: "instruction", Value: "local.get", Index: 78},
		{Type: "number", Value: "0", Index: 88},
		{Type: "instruction", Value: "local.get", Index: 95},
		{Type: "number", Value: "1", Index: 105},
		{Type: "instruction", Value: "i32.mul", Index: 112},
	}

	tokens := Tokenize(input)

	if len(tokens) != len(expectedOutput) {
		t.Errorf("Expected %d, got %d", len(expectedOutput), len(tokens))
		return
	}

	for i, token := range tokens {
		if token.Value != expectedOutput[i].Value {
			t.Errorf("Expected token value %s, got %s", expectedOutput[i].Value, token.Value)
		}
	}
}

// Test Const
func TestConst(t *testing.T) {
	input := `(module
				(func (export "operationWithInternalVariable") (param i32 i32) (result i32)
					local.get 0
					i32.const 10
					i32.add)
				)
			`
	expectedOutput := []types.Token{
		{Type: "token", Value: "module", Index: 1},
		{Type: "token", Value: "func", Index: 13},
		{Type: "token", Value: "export", Index: 19},
		{Type: "literal", Value: "\"operationWithInternalVariable\"", Index: 26},
		{Type: "token", Value: "param", Index: 60},
		{Type: "typeNum", Value: "i32", Index: 66},
		{Type: "typeNum", Value: "i32", Index: 70},
		{Type: "token", Value: "result", Index: 76},
		{Type: "typeNum", Value: "i32", Index: 83},
		{Type: "instruction", Value: "local.get", Index: 93},
		{Type: "number", Value: "0", Index: 103},
		{Type: "instruction", Value: "i32.const", Index: 110},
		{Type: "number", Value: "10", Index: 120},
		{Type: "instruction", Value: "i32.add", Index: 128},
	}

	tokens := Tokenize(input)

	if len(tokens) != len(expectedOutput) {
		t.Errorf("Expected %d, got %d", len(expectedOutput), len(tokens))
		return
	}

	for i, token := range tokens {
		if token.Value != expectedOutput[i].Value {
			t.Errorf("Expected token value %s, got %s", expectedOutput[i].Value, token.Value)
		}
	}
}

// Test Divisions
func TestDivisions(t *testing.T) {
	input := `(module
				(func (export "divNumbers") (param i32 i32) (result i32)
					local.get 0
					local.get 1
					i32.div)
				)`
	expectedOutput := []types.Token{
		{Type: "token", Value: "module", Index: 5},
		{Type: "token", Value: "func", Index: 17},
		{Type: "token", Value: "export", Index: 23},
		{Type: "literal", Value: "\"divNumbers\"", Index: 30},
		{Type: "token", Value: "param", Index: 45},
		{Type: "typeNum", Value: "i32", Index: 51},
		{Type: "typeNum", Value: "i32", Index: 55},
		{Type: "token", Value: "result", Index: 61},
		{Type: "typeNum", Value: "i32", Index: 68},
		{Type: "instruction", Value: "local.get", Index: 78},
		{Type: "number", Value: "0", Index: 88},
		{Type: "instruction", Value: "local.get", Index: 95},
		{Type: "number", Value: "1", Index: 105},
		{Type: "instruction", Value: "i32.div", Index: 112},
	}

	tokens := Tokenize(input)

	if len(tokens) != len(expectedOutput) {
		t.Errorf("Expected %d, got %d", len(expectedOutput), len(tokens))
		return
	}

	for i, token := range tokens {
		if token.Value != expectedOutput[i].Value {
			t.Errorf("Expected token value %s, got %s", expectedOutput[i].Value, token.Value)
		}
	}
}
