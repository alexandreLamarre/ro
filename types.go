package ro

import "iter"

func empty[E any]() iter.Seq[E] {
	return func(yield func(E) bool) {}
}

type intType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
