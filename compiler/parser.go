package compiler

import (
	"log"
	"luna/types"
	"strconv"
)

// Parsing guarantees that the input program is syntactically correct,
// e.g. it is properly constructed,
// but it doesn’t guarantee a successful execution
// Runtime errors may still be present

// This parser is a Go implementation of the Chasm parser
// See https://blog.scottlogic.com/2019/05/17/webassembly-compiler.html#the-parser

// Building an iterator emulator
// I tried to emulate the Javascript's array[Symbol.iterator]()
type iteratorEmulator int
type iteratorEmulatorStruct struct {
	token types.Token
	done  bool
}

func (i *iteratorEmulator) next(arr []types.Token) iteratorEmulatorStruct {
	token := arr[int(*i)]
	*i += 1
	if int(*i) >= len(arr) {
		return iteratorEmulatorStruct{
			token: token,
			done:  true,
		}
	}
	return iteratorEmulatorStruct{
		token: token,
		done:  false,
	}
}

// The parse receives the array of Tokens and creates an AST (abstract syntax tree)
// See - https://en.wikipedia.org/wiki/Abstract_syntax_tree
func Parser(tokens []types.Token) []types.AstNode {
	if len(tokens) == 0 {
		log.Fatal("No token to parse")
	}

	nodes := []types.AstNode{}
	index := 0
	// iterator
	iterator := iteratorEmulator(0)

	currentToken := iterator.next(tokens)

	if currentToken.done {
		return nodes
	}

	nextToken := iterator.next(tokens)

	eatToken := func(tokenVal string) {
		isValid, _ := strconv.ParseBool(tokenVal)
		if isValid && tokenVal != currentToken.token.Value {
			log.Fatal(`Unexpected token`)
		}

		currentToken = nextToken
		if currentToken.done {
			return
		}
		nextToken = iterator.next(tokens)
	}

	for index < len(tokens) {
		nodes = append(nodes, parseStatement(&currentToken, eatToken))
		index++
	}

	return nodes
}

func parseStatement(currentToken *iteratorEmulatorStruct, eatToken func(val string)) types.AstNode {
	if currentToken.token.Type == "keyword" {
		switch currentToken.token.Value {
		// We parse the keyword "log"
		case "log":
			eatToken("")

			return types.AstNode{
				Type: "logStatement",
				// log will print something so we evaluate what comes after that
				Expression: parseExpression(currentToken, eatToken),
			}
		}
	}
	return types.AstNode{}
}

func parseExpression(currentToken *iteratorEmulatorStruct, eatToken func(val string)) types.ExpressionNode {
	switch currentToken.token.Type {
	case "number":
		number, _ := strconv.Atoi(currentToken.token.Value)
		eatToken("")
		return types.ExpressionNode{
			Type:  "numberLiteral",
			Value: number,
		}
	}

	return types.ExpressionNode{}
}
