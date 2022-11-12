package compiler

import (
	"log"
	"luna/defaults"
	"luna/texts"
	"luna/types"
	"strconv"
)

// A WebAssembly module is organized into sections

// Each section consists of:
// - a one-byte section id,
// - the  size of the contents, in bytes,
// - the actual contents, whose structure is depended on the section id (which for most case encodes a Vector)

// Every section is optional, any omitted section will be treated as the section being present
// with empty contents.
// So our binary should resemble something like this

// 0000000: 0061 736d                                 ; WASM_BINARY_MAGIC
// 0000004: 0100 0000                                 ; WASM_BINARY_VERSION
// ; section "Type" (1)
// 0000008: 01                                        ; section code
// 0000009: 00                                        ; section size (guess)
// 000000a: 01                                        ; num types
// ; func type 0
// 000000b: 60                                        ; func
// 000000c: 02                                        ; num params
// 000000d: 7f                                        ; i32
// 000000e: 7f                                        ; i32
// 000000f: 01                                        ; num results
// 0000010: 7f                                        ; i32
// 0000009: 07                                        ; FIXUP section size
// ; section "Function" (3)
// 0000011: 03                                        ; section code
// 0000012: 00                                        ; section size (guess)
// 0000013: 01                                        ; num functions
// 0000014: 00                                        ; function 0 signature index
// 0000012: 02                                        ; FIXUP section size
// ; section "Export" (7)
// 0000015: 07                                        ; section code
// 0000016: 00                                        ; section size (guess)
// 0000017: 01                                        ; num exports
// 0000018: 06                                        ; string length
// 0000019: 6164 6454 776f                           addTwo  ; export name
// 000001f: 00                                        ; export kind
// 0000020: 00                                        ; export func index
// 0000016: 0a                                        ; FIXUP section size
// ; section "Code" (10)
// 0000021: 0a                                        ; section code
// 0000022: 00                                        ; section size (guess)
// 0000023: 01                                        ; num functions
// ; function body 0
// 0000024: 00                                        ; func body size (guess)
// 0000025: 00                                        ; local decl count
// 0000026: 20                                        ; local.get
// 0000027: 00                                        ; local index
// 0000028: 20                                        ; local.get
// 0000029: 01                                        ; local index
// 000002a: 6a                                        ; i32.add
// 000002b: 0b                                        ; end
// 0000024: 07                                        ; FIXUP func body size
// 0000022: 09                                        ; FIXUP section size
// ; section "name"
// 000002c: 00                                        ; section code
// 000002d: 00                                        ; section size (guess)
// 000002e: 04                                        ; string length
// 000002f: 6e61 6d65                                name  ; custom section name
// 0000033: 02                                        ; local name type
// 0000034: 00                                        ; subsection size (guess)
// 0000035: 01                                        ; num functions
// 0000036: 00                                        ; function index
// 0000037: 00                                        ; num locals
// 0000034: 03                                        ; FIXUP subsection size
// 000002d: 0a                                        ; FIXUP section size

// For a way more detailed explanation
// See https://webassembly.github.io/spec/core/binary/modules.html

type sectionData []interface{}
type Module []interface{}

// Recursively flatten nested arrays/slices
// Tried to emulate Javascript's Array.prototype.flat()
func flatten(input sectionData) sectionData {
	output := sectionData{}
	recursively(0, input, &output)
	return output
}

func recursively(index int, input sectionData, output *sectionData) {
	if index >= len(input) {
		return
	}
	value, ok := input[index].(sectionData)
	if ok {
		recursively(0, value, output)
	} else {
		*output = append((*output), input[index])
	}

	recursively(index+1, input, output)
}

func createSection(secType interface{}, data sectionData) sectionData {
	encodeData := encodeVector(data)

	section := sectionData{}
	section = append(section, secType)
	section = append(section, encodeData...)
	return section
}

// Encode vectors
func encodeVector(data sectionData) []interface{} {
	encoded := EncodeUnsignedLEB128(uint(len(data)))
	newEncode := sectionData{}

	for _, v := range encoded {
		newEncode = append(newEncode, interface{}(v))
	}
	vector := sectionData{}

	vector = append(vector, newEncode...)
	vector = append(vector, flatten(data)...)
	return vector
}

func omologateEncoded(num uint) sectionData {
	encodedLocalZero := EncodeUnsignedLEB128(num)
	tmp := sectionData{}
	for _, v := range encodedLocalZero {
		tmp = append(tmp, interface{}(v))
	}
	return tmp
}

// So let's start building our compiler
func Compile(ast []types.AstNode) Module {

	var module = Module{}
	// The final module array should resemble
	// [
	// 	MAGIC,
	// 	VERSION,
	//	SECTION_TYPE (1),
	// 	FUNCTION_TYPE (0),
	//	SECTION_FUNCTION (3),
	// 	SECTION_EXPORT (7),
	// 	SECTION_CODE (10),
	// 	FUNCTION_BODY (0),
	// ]

	var SECTION_TYPE = sectionData{}
	var SECTION_FUNCTION = sectionData{}
	var SECTION_EXPORT = sectionData{}
	var SECTION_CODE = sectionData{}

	var functionNumbers = []uint{}

	functionType := sectionData{}
	exportData := sectionData{}
	functionBody := sectionData{}
	code := sectionData{}

	for index, node := range ast {
		switch node.Type {
		case texts.ModuleStatement:
			// MAGIC and VERSION don't change until a newer version of WebAssembly gets released
			module = append(module, defaults.MAGIC...)
			module = append(module, defaults.VERSION...)
		case texts.FuncStatement:
			// num of types (i32, f32, i64, f64) inside the function
			functionType = append(functionType, sectionData{0x01}...)

			// Function types classify the signature of functions, mapping a vector of parameters to a vector of results.
			functionType = append(functionType, sectionData{types.FuncType}...)

			// Keep track of functions number and indexes
			if len(functionNumbers) == 0 {
				functionNumbers = append(functionNumbers, 0x00)
				break
			}

			functionIndex := functionNumbers[len(functionNumbers)-1] + 1
			functionNumbers = append(functionNumbers, functionIndex)

		case texts.ParamStatement:
			paramsAlreadyAdded := false

			value, ok := node.Expression.Value.(string)
			if !ok {
				log.Fatal("not a string")
			}
			params := sectionData{types.ValType[value]}
			// Params
			// Get num of params dynamically
			for idx, _node := range ast[index+1:] {
				nextNode := ast[index+(idx+1)]

				if (_node.Type != texts.ParamStatement) && (nextNode.Type != texts.TypeNum32) {
					// All the parameters were parsed already
					if idx == 0 {
						paramsAlreadyAdded = true
					}
					break
				}

				params = append(params, sectionData{types.ValType[value]})
			}

			if paramsAlreadyAdded {
				continue
			}

			functionType = append(functionType, encodeVector(params)...)

		case texts.ResultStatement:
			value, ok := node.Expression.Value.(string)
			if !ok {
				log.Fatal("not a string")
			}
			// Result
			// Currently WebAssembly supports only one returned result
			functionType = append(functionType, encodeVector(sectionData{
				types.ValType[value],
			})...)

		// Instructions
		// https://webassembly.github.io/spec/core/binary/instructions.html
		case texts.GetLocalInstruction:
			var localIndex uint
			value, ok := node.Expression.Value.(string)
			if !ok {
				log.Fatal("Not string")
			}

			v, _ := strconv.Atoi(value)
			localIndex = uint(v)

			code = append(code, sectionData{node.MapTo}...)
			code = append(code, omologateEncoded(localIndex)...)

		case texts.CallStatement:
			var _index uint = 0
			code = append(code, sectionData{defaults.Opcodes["call"]}...)
			for _, n := range ast {

				if n.Type == texts.FuncStatement {
					if n.Expression.Value == node.Expression.Value {
						code = append(code, omologateEncoded(_index)...)
					}
					_index++
				}
			}

		case texts.InternalInstruction:
			var localIndex uint
			value, ok := node.Expression.Value.(string)
			if !ok {
				log.Fatal("Not string")
			}

			v, _ := strconv.Atoi(value)
			localIndex = uint(v)

			code = append(code, sectionData{node.MapTo}...)
			code = append(code, omologateEncoded(localIndex)...)

		// Export section
		case texts.ExportStatement:
			value, ok := node.Expression.Value.(string)
			if !ok {
				log.Fatal("not a string")
			}

			// Let's build export section
			encodedString := encodeString(value)

			// number of exports
			// We only have one export so the number is one
			exportData = append(exportData, sectionData{0x01}...)
			exportData = append(exportData, encodeVector(encodedString)...)
			exportData = append(exportData, sectionData{defaults.ExportSection["func"]}...)
			// Export type index
			// We only have one exported function so the index is 0
			exportData = append(exportData, sectionData{0x00}...)

		// Remember the concept of Stack Machine
		case texts.FuncInstruction:
			code = append(code, sectionData{node.MapTo}...)

			// Put all the section code together
			functionBodyData := sectionData{}

			// Locals declaration count
			// See https://webassembly.github.io/spec/core/binary/modules.html#code-section:~:text=Local%20declarations
			functionBodyData = append(functionBodyData, sectionData{0x00}...)
			functionBodyData = append(functionBodyData, code...)
			functionBodyData = append(functionBodyData, sectionData{defaults.Opcodes["end"]})

			// Number of functions
			functionBody = append(functionBody, sectionData{0x01})
			functionBody = append(functionBody, encodeVector(functionBodyData)...)

		}

	}

	// Type Section
	// The type section has the id 1. It decodes into a vector of function types that represent the  component of a module.
	// See https://webassembly.github.io/spec/core/binary/modules.html#type-section
	SECTION_TYPE = createSection(defaults.Section["type"], functionType)
	// Func Section
	// The function section has the id 3. It decodes into a vector of type indices that represent the type fields
	// of the functions in the funcs component of a module.
	functionIndexes := sectionData{}
	for _, number := range functionNumbers {
		functionIndexes = append(functionIndexes, number)
	}
	// See https://webassembly.github.io/spec/core/binary/modules.html#function-section
	SECTION_FUNCTION = createSection(defaults.Section["func"], encodeVector(functionIndexes))
	// Code section
	// The code section has the id 10. It decodes into a vector of code entries that are pairs of value type vectors and expressions.
	// See https://webassembly.github.io/spec/core/binary/modules.html#code-section
	SECTION_CODE = createSection(defaults.Section["code"], functionBody)
	// Export Section
	// The export section has the id 7.
	// It decodes into a vector of exports that represent the  component of a module.
	// See https://webassembly.github.io/spec/core/binary/modules.html#export-section
	SECTION_EXPORT = createSection(defaults.Section["export"], exportData)

	module = append(module, SECTION_TYPE...)
	module = append(module, SECTION_FUNCTION...)
	module = append(module, SECTION_EXPORT...)
	module = append(module, SECTION_CODE...)

	return module
}
