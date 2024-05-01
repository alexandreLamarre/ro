//go:build goexperiment.rangefunc

package ro

import "iter"

// Count returns an infinite iterator starting from start and incrementing by step
func Count[T intType](start, step T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := start; ; i += step {
			if !yield(i) {
				break
			}
		}
	}
}

// Repeat returns an infinite iterator that yields elem indefinitely
func Repeat[T intType](elem T) iter.Seq[T] {
	return Count(elem, 0)
}

// CycleSlice returns an infinite iterator that cycles repeatedly through the elements of the slice
func CycleSlice[T any](arr []T) iter.Seq[T] {
	done := false
	if len(arr) == 0 {
		return empty[T]()
	}
	return func(yield func(T) bool) {
		for {
			for _, v := range arr {
				if !yield(v) {
					done = true
					break
				}
			}
			if done {
				break
			}
		}
	}
}

// Cycle returns an infinite iterator that cycles repeatedly through the elements of sequence
func Cycle[T any](seq iter.Seq[T]) iter.Seq[T] {
	done := false
	return func(yield func(T) bool) {
		for {
			for v := range seq {
				if !yield(v) {
					done = true
					break
				}
			}
			if done {
				break
			}
		}
	}
}
