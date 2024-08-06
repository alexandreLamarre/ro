package ro

import "iter"

func empty[T any]() iter.Seq[T] {
	return func(_ func(T) bool) {}
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

// ToSlice is a convenience wrapper to convert an iterator to a slice
//
// This will block until the iterator is exhausted.
func ToSlice[T any](seq iter.Seq[T]) []T {
	res := []T{}
	for v := range seq {
		res = append(res, v)
	}
	return res
}

// FromSlice is a convenience wrapper to convert a slice to an iterator
func FromSlice[T any](arr []T) iter.Seq[T] {
	if len(arr) == 0 {
		return empty[T]()
	}
	return func(yield func(T) bool) {
		for _, v := range arr {
			if !yield(v) {
				break
			}
		}
	}
}

// FromString is a convenience wrapper to convert a string to an iterator
func FromString(s string) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, r := range s {
			if !yield(r) {
				break
			}
		}
	}
}

// Extend pads an iterator with an empty struct to conform to an iter.Seq2 type
func Extend[T any](seq iter.Seq[T]) iter.Seq2[struct{}, T] {
	return func(yield func(struct{}, T) bool) {
		for v := range seq {
			if !yield(struct{}{}, v) {
				break
			}
		}
	}
}
