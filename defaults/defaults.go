package defaults

// Module magic \asm and version
// URL: https://webassembly.github.io/spec/core/binary/modules.html#binary-version
var (
	MAGIC   = []interface{}{0x00, 0x61, 0x73, 0x6d}
	VERSION = []interface{}{0x01, 0x00, 0x00, 0x00}
)

// Opcodes
// URL on https://webassembly.github.io/spec/core/binary/instructions.html
var (
	// unreachable = 0x00
	block       = 0x02
	loop        = 0x03
	br          = 0x0c
	br_if       = 0x0d
	end         = 0x0b
	call        = 0x10
	get_local   = 0x20
	set_local   = 0x21
	i32_store_8 = 0x3a
	i32_const   = 0x41
	f32_const   = 0x43
	i32_eqz     = 0x45
	i32_eq      = 0x46
	f32_eq      = 0x5b
	f32_lt      = 0x5d
	f32_gt      = 0x5e
	i32_and     = 0x71
	f32_add     = 0x92
	f32_sub     = 0x93
	f32_mul     = 0x94
	f32_div     = 0x95
)

var Opcodes = map[string]interface{}{
	"block":       block,
	"loop":        loop,
	"br":          br,
	"br_if":       br_if,
	"end":         end,
	"call":        call,
	"get_local":   get_local,
	"set_local":   set_local,
	"i32_store_8": i32_store_8,
	"i32_const":   i32_const,
	"f32_const":   f32_const,
	"i32_eqz":     i32_eqz,
	"i32_eq":      i32_eq,
	"f32_eq":      f32_eq,
	"f32_lt":      f32_lt,
	"f32_gt":      f32_gt,
	"i32_and":     i32_and,
	"f32_add":     f32_add,
	"f32_sub":     f32_sub,
	"f32_mul":     f32_mul,
	"f32_div":     f32_div,
}

var OperatorsOpcodes = map[string]interface{}{
	"+":  f32_add,
	"-":  f32_sub,
	"*":  f32_mul,
	"/":  f32_div,
	"==": f32_eq,
	">":  f32_gt,
	"<":  f32_lt,
	"&&": i32_and,
}

// Export section
// Based on http://webassembly.github.io/spec/core/binary/modules.html#export-section
var ExportSection = map[string]interface{}{
	"funx":   0x00,
	"table":  0x01,
	"mem":    0x02,
	"global": 0x03,
}