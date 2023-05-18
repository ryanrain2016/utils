package utils

type UInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type SInt interface {
	int | int8 | int16 | int32 | int64 | UInt
}

type Int interface {
	SInt | UInt
}

type Float interface {
	float32 | float64
}

type OrderedNumber interface {
	Int | Float
}

type Complex interface {
	complex64 | complex128
}

type Number interface {
	OrderedNumber | Complex
}

type Ordered interface {
	OrderedNumber | string
}

type Addable interface {
	Number | string
}
