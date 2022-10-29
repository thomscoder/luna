package compiler

import (
	"fmt"
	"log"
	"luna/types"
	"strconv"
)

// Building an iterator emulator
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
			token: types.Token{},
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
		log.Fatal("No tokens to parse")
	}
	nodes := []types.AstNode{}
	index := 0
	// iterator
	iterator := iteratorEmulator(0)

	currentToken := iterator.next(tokens)
	var nextToken iteratorEmulatorStruct

	fmt.Println("FUCK", nextToken)
	if currentToken.done {
		return nodes
	}

	nextToken = iterator.next(tokens)

	eatToken := func(tokenVal string) {
		isValid, _ := strconv.ParseBool(tokenVal)
		if isValid && tokenVal != currentToken.token.Value {
			log.Fatal(`Unexpected token`)
		}

		currentToken = nextToken
		nextToken = iterator.next(tokens)
	}

	for index < len(tokens) {
		nodes = append(nodes, parseStatement(currentToken, eatToken))
		index++
	}
	return nodes
}

func parseStatement(currentToken iteratorEmulatorStruct, eatToken func(val string)) types.AstNode {
	if currentToken.token.Type == "keyword" {
		switch currentToken.token.Value {
		case "log":
			eatToken("")
			return types.AstNode{
				Type:       "logStatement",
				Expression: parseExpression(currentToken, eatToken),
			}
		}
	}
	return types.AstNode{}
}

func parseExpression(currentToken iteratorEmulatorStruct, eatToken func(val string)) types.ExpressionNode {
	var node types.ExpressionNode
	switch currentToken.token.Type {
	case "number":
		number, _ := strconv.Atoi(currentToken.token.Value)
		node.Type = "numberLiteral"
		node.Value = number
		eatToken("")
		return node
	}

	return types.ExpressionNode{}
}
