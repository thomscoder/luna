package compiler

import (
	"luna/types"
	"regexp"
	"strings"
)

var keywords = []string{
	"log",
}

// Higher order function
func match(regex string, whichType string) func(string, int) types.Matcher {

	return func(input string, index int) types.Matcher {
		substr := input[index:]
		rxp := regexp.MustCompile(regex)
		match := rxp.FindAllString(substr, -1)

		if len(match) > 0 {
			return types.Matcher{Type: whichType, Value: match[0]}
		}

		return types.Matcher{}
	}
}

var matchers = []func(string, int) types.Matcher{
	match("^[.0-9]+", "number"),
	match("^("+strings.Join(keywords, "|")+")", "keyword"),
	match("^\\s+", "whitespace"),
}

func Tokenize(input string) interface{} {
	tokens := []types.Token{}
	matches := []types.Matcher{}
	index := 0
	for index < len(input) {
		for _, m := range matchers {
			matchFound := m(input, index)
			// Prevent panic if no match is found
			if (matchFound == types.Matcher{}) {
				continue
			}
			matches = append(matches, matchFound)
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
