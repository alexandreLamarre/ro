package ro

import "iter"

// Count returns an infinite iterator starting from start and incrementing by step
func Count[i intType](start, step i) iter.Seq[i] {
	return func(yield func(i) bool) {
		for i := start; ; i += step {
			if !yield(i) {
				break
			}
		}
	}
}

// Repeat returns an infinite iterator that yields elem indefinitely
func Repeat[i intType](elem i) iter.Seq[i] {
	return Count(elem, 0)
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
