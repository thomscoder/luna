package types

// Number Types
// See https://webassembly.github.io/spec/core/binary/types.html#number-types
var NumTypes = map[string]interface{}{
	"i32": 0x7f,
	"i64": 0x7e,
	"f32": 0x7d,
	"f64": 0x7c,
}

// Vector Types
// See https://webassembly.github.io/spec/core/binary/types.html#vector-types
const V128 = 0x7b

// Function Types
// See https://webassembly.github.io/spec/core/binary/types.html#function-types
const FuncType = 0x60

// Value types
// See https://webassembly.github.io/spec/core/binary/types.html#value-types
var ValType = map[string]interface{}{
	"i32": 0x7f,
	"i64": 0x7e,
	"f32": 0x7d,
	"f64": 0x7c,
}
