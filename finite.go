package ro

import "iter"

// Range returns an iterator that yields integers from start to end (exclusive) by step
func Range[i intType](start, end, step i) iter.Seq[i] {
	if step == 0 {
		return empty[i]()
	}
	return func(yield func(i) bool) {
		for i := start; i < end; i += step {
			if !yield(i) {
				break
			}
		}
	}
}
