//go:build goexperiment.rangefunc

package ro

import "iter"

func empty[E any]() iter.Seq[E] {
	return func(_ func(E) bool) {}
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

// SeqAsIter is a convenience wrapper to convert a slice to an iterator
func SeqAsIter[T any](seq []T) iter.Seq[T] {
	if len(seq) == 0 {
		return empty[T]()
	}
	return func(yield func(T) bool) {
		for _, v := range seq {
			if !yield(v) {
				break
			}
		}
	}
}

// StringAsSeq is a convenience wrapper to convert a string to an iterator
func StringAsSeq(s string) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, r := range s {
			if !yield(r) {
				break
			}
		}
	}
}
