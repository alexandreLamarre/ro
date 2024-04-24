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

// Cycle returns an infinite iterator that cycles repeatedly through the elements of seq
func Cycle[V any](seq []V) iter.Seq[V] {
	done := false
	if len(seq) == 0 {
		return empty[V]()
	}
	return func(yield func(V) bool) {
		for {
			for _, v := range seq {
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
