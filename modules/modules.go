package modules

func Emitter() []uint8 {
	arr := []uint8{}

	arr = append(arr, magic...)
	arr = append(arr, version...)

	return arr
}
