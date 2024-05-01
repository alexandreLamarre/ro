//go:build goexperiment.rangefunc

package ro

import (
	"iter"
)

// Accumulate returns an iterator that yields the accumulated sum of the elements in the slice
func Accumulate[T intType](arr []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		sum := T(0)
		for _, v := range arr {
			sum += v
			if !yield(sum) {
				break
			}
		}
	}
}

// AccumulateIter returns an iterator that yields the accumulated sum of the elements yielded by seq
func AccumulateIter[T intType](seq iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		sum := T(0)
		for v := range seq {
			sum += v
			if !yield(sum) {
				break
			}
		}
	}
}

// AccumulateFunc returns an iterator that yields the accumulated result of applying f to the elements in the slice
func AccumulateFunc[T any](arr []T, f func(T, T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		var agg T
		for _, v := range arr {
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

// Batch returns an iterator that yields batches of size elements from the slice
func Batch[T any](arr []T, size int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for i := 0; i < len(arr); i += size {
			end := min(i+size, len(arr))
			if !yield(arr[i:end]) {
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

// Chain returns an iterator that yields the elements of the slice in order
func Chain[T any](arr ...[]T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range arr {
			for _, v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// ChainIter returns an iterator that yields the elements yielded by seqs in order
func ChainIter[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, it := range seqs {
			for v := range it {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Drop returns an iterator that yields elements not matching the predicate
func Drop[T any](arr []T, predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range arr {
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

// Filter returns an iterator that yields elements matching the predicate
func Filter[T any](arr []T, predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range arr {
			if predicate(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
}

// FilterIter returns an iterator that yields elements matching the predicate
func FilterIter[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if predicate(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
}

// Pairwise returns an iterator that yields pairs of adjacent elements in the slice
// If the slice has less than 2 elements, the empty iterator is returned
func PairWise[T any](arr []T) iter.Seq[[2]T] {
	return func(yield func([2]T) bool) {
		if len(arr) < 2 {
			return
		}
		for i := 0; i < len(arr)-1; i++ {
			if !yield([2]T{arr[i], arr[i+1]}) {
				break
			}
		}
	}
}

// PairWiseIter returns an iterator that yields pairs of adjacent elements yielded by seq
// If the sequence has less than 2 elements, the empty iterator is returned
func PairWiseIter[T any](seq iter.Seq[T]) iter.Seq[[2]T] {
	return func(yield func([2]T) bool) {
		var prev T
		first := true
		for v := range seq {
			if first {
				prev = v
				first = false
				continue
			}
			if !yield([2]T{prev, v}) {
				break
			}
			prev = v
		}
	}
}

// While returns an iterator that yields elements of the slice until the predicate is false
func While[T any](arr []T, predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range arr {
			if !predicate(v) {
				break
			}
			if !yield(v) {
				break
			}
		}
	}
}

// WhileIter returns an iterator that yields elements of the sequence until the predicate is false
func WhileIter[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if !predicate(v) {
				break
			}
			if !yield(v) {
				break
			}
		}
	}
}

// Limit returns an iterator that yields the up to the first n elements of the slice
func Limit[T any](arr []T, n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i, v := range arr {
			if i >= n {
				break
			}
			if !yield(v) {
				break
			}
		}
	}
}

// LimitIter returns an iterator that yields the up to the first n elements yielded by the sequence
func LimitIter[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		k := 0
		for v := range seq {
			if k >= n {
				break
			}
			k++
			if !yield(v) {
				break
			}
		}
	}
}

// Apply returns an iterator that yields the result of applying f to each element in the slice
func Apply[U, V any](arr []U, f func(U) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range arr {
			if !yield(f(v)) {
				break
			}
		}
	}
}

// ApplyIter returns an iterator that yields the result of applying f to each element yielded by the sequence
func ApplyIter[U, V any](seq iter.Seq[U], f func(U) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !yield(f(v)) {
				break
			}
		}
	}
}

// Tee returns n iterators that yield the elements of the sequence
// If n == 1, the only element in the slice will be the original seq
func Tee[T any](seq iter.Seq[T], n int) []iter.Seq[T] {
	if n == 1 {
		return []iter.Seq[T]{seq}
	}
	res := []iter.Seq[T]{}
	for i := 0; i < n; i++ {
		res = append(res, func(yield func(T) bool) {
			for v := range seq {
				if !yield(v) {
					break
				}
			}
		})
	}
	return res
}
