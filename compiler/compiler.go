// Following the ../defaultss/main.wat
// This is the Hex dump of the above cited Wasm Binary version

// WASM MAGIC
// 0x00, // \0
// 0x61, // a
// 0x73, // s
// 0x6d, // m

//WASM VERSION
// 0x01, // 1
// 0x00, // 0
// 0x00, // 0
// 0x00, // 0

// SECTION "Type"
// 0x01, // section code
// 0x07, // section size
// 0x01, // num types

// type 0
// 0x60, // func
// 0x02, // num params
// 0x7f, // i32
// 0x7f, // i32
// 0x01, // num results
// 0x7f, // i32

// SECTION "Function"
// 0x03, // section code
// 0x02, // section size
// 0x01, // num functions
// 0x00, // function 0 signature index

// section "Export"
// 0x07, // section export
// 0x07, // section size
// 0x01, // num exports
// 0x03, // string length

// "add" EXPORT NAME
// 0x61, // a
// 0x64, // d
// 0x64, // d
// 0x00, // 0

// EXPORT KIND
// 0x00, // export func index

// SECTION "Code"
// 0x0a, // section code
// 0x09, // section size
// 0x01, // num function

// FUNCTION BODY
// 0x07, // func body size
// 0x00, // local decl count
// 0x20, // local.get
// 0x00, // local index
// 0x20, // local.get
// 0x01, // local index
// 0x6a, // i32.add
// 0x0b, // end

// For a detailed explanation of sections
// See https://webassembly.github.io/spec/core/binary/modules.html#sections

// Luckily our language is not that robust yet, so we can keep it simpler

package compiler

import (
	"luna/defaults"
	"luna/types"
)

func appendToCode(code *[]int, val []interface{}) {

	for i := range val {
		value, ok := val[i].(int)

		if !ok {

			value = int(value)
		}

		(*code)[i] = value
	}
}

func computeArray(origArr *[4]byte) []interface{} {
	arr := []interface{}{}
	for _, v := range *origArr {
		arr = append(arr, v)
	}
	return arr
}

// The emitter takes the Ast (see ./parser.go)
// and emits proper instructions
func Emitter(ast []types.AstNode) []int {
	code := make([]int, 29)
	// Magic header and Version need to be set for all Wasm modules
	// They remain the same until new Version of Web Assembly gets released
	appendToCode(&code, defaults.MAGIC)
	appendToCode(&code, defaults.VERSION)

	emitExpression := func(node types.ExpressionNode) {
		switch node.Type {
		case "numberLiteral":
			val, ok := node.Value.(int)
			if !ok {
				val = int(val)
			}

			appendToCode(&code, []interface{}{defaults.Opcodes["f32_const"]})
			arr := Ieee754(val)
			appendToCode(&code, computeArray(&arr))
		}
	}

	for _, node := range ast {
		switch node.Type {
		case "logStatement":
			emitExpression(node.Expression)
			appendToCode(&code, []interface{}{defaults.Opcodes["call"]})

			appendToCode(&code, []interface{}{EncodeUnsignedLEB128(0)})
		}
	}

	return code
}
