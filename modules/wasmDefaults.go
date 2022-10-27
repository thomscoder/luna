package modules

// Based on https://webassembly.github.io/spec/core/binary/modules.html#binary-version
var (
	magic   = []uint8{0x00, 0x61, 0x73, 0x6d}
	version = []uint8{0x01, 0x00, 0x00, 0x00}
)
