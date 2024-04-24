package ro

import "iter"

// Count returns an infinite iterator starting from start and incrementing by step
func Count(start, step int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; ; i += step {
			if !yield(i) {
				break
			}
		}
	}
}
