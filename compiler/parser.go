package compiler

import (
	"log"
	"luna/defaults"
	"luna/texts"
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
		nodes = append(nodes, parseStatement(&currentToken, eatToken, &index))
		index++
	}

	return nodes
}

func parseStatement(currentToken *iteratorEmulatorStruct, eatToken func(val string), index *int) types.AstNode {
	if currentToken.token.Type == texts.TypeToken {
		switch currentToken.token.Value {
		case "module":
			eatToken("module")
			return types.AstNode{
				Type: texts.ModuleStatement,
				// log will print something so we evaluate what comes after that
				Expression: types.ExpressionNode{},
			}

		case "func":
			eatToken("func")

			return types.AstNode{
				Type: texts.FuncStatement,
				// log will print something so we evaluate what comes after that
				Expression: types.ExpressionNode{},
			}
		case "export":
			eatToken("export")

			return types.AstNode{
				Type: texts.ExportStatement,
				// log will print something so we evaluate what comes after that
				Expression: parseExpression(currentToken, eatToken, index),
			}
		case "result":
			eatToken("result")

			return types.AstNode{
				Type: texts.ResultStatement,
				// log will print something so we evaluate what comes after that
				Expression: parseExpression(currentToken, eatToken, index),
			}
		case "param":
			eatToken("param")

			return types.AstNode{
				Type: texts.ParamStatement,
				// log will print something so we evaluate what comes after that
				Expression: types.ExpressionNode{},
			}
		}
	}

	if currentToken.token.Type == texts.TypeInstruction {
		switch currentToken.token.Value {
		case "local.get":
			eatToken("local.get")
			return types.AstNode{
				Type:       texts.GetLocalInstruction,
				Expression: parseExpression(currentToken, eatToken, index),
				MapTo:      defaults.Opcodes["get_local"],
			}
		}
		switch currentToken.token.Value {
		case "i32.add":
			eatToken("i32.add")
			return types.AstNode{
				Type:       texts.FuncInstruction,
				Expression: types.ExpressionNode{},
				MapTo:      defaults.Opcodes["i32_add"],
			}
		}
	}

	if currentToken.token.Type == texts.TypeNum {
		switch currentToken.token.Value {
		case "i32":
			eatToken("i32")
			return types.AstNode{
				Type:       texts.TypeNum32,
				Expression: types.ExpressionNode{},
				MapTo:      types.ValType["i32"],
			}
		}
	}

	return types.AstNode{}
}

func parseExpression(currentToken *iteratorEmulatorStruct, eatToken func(val string), index *int) types.ExpressionNode {
	switch currentToken.token.Type {
	case "number":
		log := types.ExpressionNode{
			Type:  texts.NumberLiteral,
			Value: currentToken.token.Value,
		}
		eatToken("number")
		*index++

		return log

	case "literal":
		log := types.ExpressionNode{
			Type:  texts.TypeLiteral,
			Value: currentToken.token.Value,
		}
		eatToken("literal")
		*index++

		return log

	case "typeNum":
		log := types.ExpressionNode{
			Type:  texts.TypeNum32,
			Value: currentToken.token.Value,
		}
		eatToken("typeNum")
		*index++

		return log

	}

	return types.ExpressionNode{}
}
