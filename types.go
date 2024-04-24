package ro

import "iter"

func empty[E any]() iter.Seq[E] {
	return func(yield func(E) bool) {}
}

type intType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type floatType interface {
	~float32 | ~float64
}

type complexType interface {
	~complex64 | ~complex128
}

type numberType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	~float32 | ~float64 | ~complex64 | ~complex128
}
