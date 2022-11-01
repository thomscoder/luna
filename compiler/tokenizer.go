package compiler

import (
	"errors"
	"luna/texts"
	"luna/types"
	"regexp"
	"strings"
)

// Inside this wasm module there are two main things:
// - A function (identified with the keyword func)
// - The export (identified with the keyword export)

// More generally speaking this module consists of 3 parts
// - Tokens: special keywords reserved by the language (e.g. func, param, module, local.get etc...)
// - Identifiers: what can be set to arbitrary values (e.g. $firstNumber, $secondNumber)
// - Value Types: defined by the Web Assembly specifications (e.g. i32)

// We define the tokens
// Tokens: special tokens reserved by the language (e.g. log)
var tokens = []string{
	"func",
	"module",
	"export",
	"result",
	"param",
}

// Hard coding identifiers for simplicity reasons (since the example won't change)
// A better and more robust regex could be implemented
var instructions = []string{
	"local\\.get",
	"i32\\.add",
}

var numTypes = []string{
	"i32",
	"i64",
	"f32",
	"f64",
}

// Bunch of names for function export
// So we do not need to change this everytime
// Might implement a regex to match anything between quotes
var literals = []string{
	"\"([^\"]+)\"",
}

// The tokenizer goes through the input (string) and gets all the matching patterns
// that represent the tokens

// Higher order function
func matchChecker(regex string, whichType string) func(string, int) (types.Matcher, error) {

	return func(input string, index int) (types.Matcher, error) {

		substr := input[index:]
		rxp := regexp.MustCompile(regex)
		match := rxp.FindString(substr)

		if len(match) > 0 {
			return types.Matcher{Type: whichType, Value: match}, nil
		}

		return types.Matcher{}, errors.New("no match found")
	}
}

var matchers = []func(string, int) (types.Matcher, error){
	matchChecker("^("+strings.Join(tokens, "|")+")", texts.TypeToken),
	matchChecker("^("+strings.Join(instructions, "|")+")", texts.TypeInstruction),
	matchChecker("^("+strings.Join(numTypes, "|")+")", texts.TypeNum),
	matchChecker("^("+strings.Join(literals, "|")+")", texts.TypeLiteral),
	matchChecker("^[0-9]+", texts.Number),
	matchChecker("^\\s+", texts.Whitespace),
}

func Tokenize(input string) []types.Token {
	tokens := []types.Token{}
	matches := []types.Matcher{}
	index := 0
	for index < len(input) {
		for _, m := range matchers {
			matchFound, notFound := m(input, index)

			// Prevent panic if no match is found
			if notFound != nil {
				continue
			}

			matches = append(matches, matchFound)
		}
		if len(matches) == 0 {
			index++
			continue
		}
		match := types.Token{
			Type:  matches[0].Type,
			Value: matches[0].Value,
			Index: index,
		}

		if match.Type != "whitespace" {
			tokens = append(tokens, match)
		}

		index += len(matches[0].Value)
		matches = []types.Matcher{}
	}
	return tokens
}
