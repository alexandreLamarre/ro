//go:build goexperiment.rangefunc

package ro

import (
	"iter"
)

// Accumulate returns an iterator that yields the accumulated sum of the elements in seq
func Accumulate[i intType](seq []i) iter.Seq[i] {
	return func(yield func(i) bool) {
		sum := i(0)
		for _, v := range seq {
			sum += v
			if !yield(sum) {
				break
			}
		}
	}
}

// AccumulateIter returns an iterator that yields the accumulated sum of the elements yielded by seq
func AccumulateIter[i intType](seq iter.Seq[i]) iter.Seq[i] {
	return func(yield func(i) bool) {
		sum := i(0)
		for v := range seq {
			sum += v
			if !yield(sum) {
				break
			}
		}
	}
}

// AccumulateFunc returns an iterator that yields the accumulated result of applying f to the elements in seq
func AccumulateFunc[T any](seq []T, f func(T, T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		var agg T
		for _, v := range seq {
			agg = f(agg, v)
			if !yield(agg) {
				break
			}
		}
	}
}

// AccumulateIterFunc returns an iterator that yields the accumulated result of applying f to the elements yielded by seq
func AccumulateIterFunc[T any](seq iter.Seq[T], f func(T, T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		var agg T
		for v := range seq {
			agg = f(agg, v)
			if !yield(agg) {
				break
			}
		}
	}
}

// Batch returns an iterator that yields batches of size elements from seq
func Batch[T any](seq []T, size int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for i := 0; i < len(seq); i += size {
			end := min(i+size, len(seq))
			if !yield(seq[i:end]) {
				break
			}
		}
	}
}

// BatchIter returns an iterator that yields batches of size elements yielded by seq
func BatchIter[T any](seq iter.Seq[T], size int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		batch := []T{}
		for v := range seq {
			batch = append(batch, v)
			if len(batch) == size {
				if !yield(batch) {
					break
				}
				batch = []T{}
			}
		}
		if len(batch) > 0 {
			yield(batch)
		}
	}
}

// Chain returns an iterator that yields the elements of seqs in order
func Chain[T any](seqs ...[]T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			for _, v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// ChainIter returns an iterator that yields the elements yielded by seqs in order
func ChainIter[T any](iterators ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, it := range iterators {
			for v := range it {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Drop returns an iterator that yields elements not matching the predicate
func Drop[T any](seq []T, predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range seq {
			if !predicate(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
}

// DropIter returns an iterator that yields elements not matching the predicate
func DropIter[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if !predicate(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
}
