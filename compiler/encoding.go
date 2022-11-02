package compiler

// According to WebAssembly specification (https://webassembly.github.io/spec/core/_download/WebAssembly.pdf)
// strings are encoded using UTF-8 encoding
// Go uses it by default
func encodeString(str string) []interface{} {
	encodedString := []interface{}{}

	for _, v := range str {
		encodedString = append(encodedString, interface{}(rune(v)))
	}

	return encodedString
}

// According to WebAssembly specification (https://webassembly.github.io/spec/core/_download/WebAssembly.pdf)
// all integers are encoded using LEB128 in either signed or unsigned variant
// See https://en.wikipedia.org/wiki/LEB128

// Implementation for Unsigned integers
// https://en.wikipedia.org/wiki/LEB128#Encode_unsigned_integer
func EncodeUnsignedLEB128(number uint) []uint {
	buff := []uint{}

	// Do while emulation
	for n := true; n; n = number != 0 {

		_byte := number & 0x7f
		number >>= 7
		if number != 0 {
			_byte |= 0x80
		}
		buff = append(buff, _byte)
	}

	return buff
}

// Implementation for the signed integers
// See javascript implementation https://en.wikipedia.org/wiki/LEB128#JavaScript_code
func EncodeSignedLEB128(number int) []int {
	number |= 0
	buff := []int{}

	for {
		_byte := number & 0x7f
		number >>= 7
		if (number == 0 && (_byte&0x40) == 0) || (number == -1 && (_byte&0x40) != 0) {
			buff = append(buff, _byte)
			return buff
		}
		buff = append(buff, _byte|0x80)
	}
}
