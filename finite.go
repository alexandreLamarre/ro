package ro

import "iter"

// Range returns an iterator that yields integers from start to end (exclusive) by step
func Range(start, end, step int) iter.Seq[int] {
	if step == 0 {
		return empty[int]()
	}
	return func(yield func(int) bool) {
		for i := start; i < end; i += step {
			if !yield(i) {
				break
			}
		}
	}
}
