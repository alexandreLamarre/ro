package ro

import "iter"

// Range returns an iterator that yields integers from start to end (exclusive) by step
func Range[T intType](start, end, step T) iter.Seq[T] {
	if step == 0 {
		return empty[T]()
	}
	return func(yield func(T) bool) {
		for i := start; i < end; i += step {
			if !yield(i) {
				break
			}
		}
	}
}
