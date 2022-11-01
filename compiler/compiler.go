package compiler

import (
	"log"
	"luna/defaults"
	"luna/texts"
	"luna/types"
	"strconv"
)

// Our binary should resemble something like this

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

// For a detailed explanation of sections
// See https://webassembly.github.io/spec/core/binary/modules.html#sections
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
	// Size of section
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

	functionType := sectionData{}
	exportData := sectionData{}
	functionBody := sectionData{}
	code := sectionData{}

	for _, node := range ast {
		switch node.Type {
		case texts.ModuleStatement:
			// MAGIC and VERSION don't change until a newer version of WebAssembly gets released
			module = append(module, defaults.MAGIC...)
			module = append(module, defaults.VERSION...)
		case texts.FuncStatement:
			// num of types
			functionType = append(functionType, sectionData{0x01}...)

			// The type section has the id 1. It decodes into a vector of function types that represent the  component of a module.
			// Function types classify the signature of functions, mapping a vector of parameters to a vector of results.
			functionType = append(functionType, sectionData{types.FuncType}...)

		case texts.ParamStatement:
			// Params
			// Num of params (harcoding 2)
			functionType = append(functionType, encodeVector(sectionData{
				types.ValType["i32"],
				types.ValType["i32"],
			})...)

		case texts.ResultStatement:
			value, ok := node.Expression.Value.(string)
			if !ok {
				log.Fatal("not a string")
			}
			// Result
			// Num of returned results (hard coding 1)
			functionType = append(functionType, encodeVector(sectionData{
				types.ValType[value],
			})...)

			//fmt.Println("fil", SECTION_EXPORT)
			// Let's build code section

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

		case texts.ExportStatement:
			value, ok := node.Expression.Value.(string)
			if !ok {
				log.Fatal("not a string")
			}

			// Let's build export section
			encodedString := encodeString(value)

			// DEPRECATED
			// The above is easier to be decode by Javascript
			// for i := 0; i < len(name); i++ {
			// 	r, _ := utf8.DecodeRuneInString(string(name[i]))
			// 	s := fmt.Sprintf("%x", r)
			// 	encodedString = append(encodedString, interface{}(s))
			// }

			// number of exports
			exportData = append(exportData, sectionData{0x01}...)
			exportData = append(exportData, encodeVector(encodedString)...)
			exportData = append(exportData, sectionData{defaults.ExportSection["func"]}...)
			exportData = append(exportData, sectionData{0x0}...)

		// Remember the concept of Stack Machine
		case texts.FuncInstruction:
			code = append(code, sectionData{node.MapTo}...)

			// Put all the section code together
			functionBodyData := sectionData{}

			functionBodyData = append(functionBodyData, sectionData{0x00}...)
			functionBodyData = append(functionBodyData, code...)
			functionBodyData = append(functionBodyData, sectionData{defaults.Opcodes["end"]})

			// num of functions
			functionBody = append(functionBody, sectionData{0x01})
			functionBody = append(functionBody, encodeVector(functionBodyData)...)

		}

	}

	// Type Section
	SECTION_TYPE = createSection(defaults.Section["type"], functionType)
	// Func Section
	SECTION_FUNCTION = createSection(defaults.Section["func"], encodeVector(sectionData{0x00}))
	// Code section
	SECTION_CODE = createSection(defaults.Section["code"], functionBody)
	// Export Section
	SECTION_EXPORT = createSection(defaults.Section["export"], exportData)

	module = append(module, SECTION_TYPE...)
	module = append(module, SECTION_FUNCTION...)
	module = append(module, SECTION_EXPORT...)
	module = append(module, SECTION_CODE...)

	return module
}
