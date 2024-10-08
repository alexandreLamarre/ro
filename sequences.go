package ro

import (
	"iter"
)

// AccumulateSlice returns an iterator that yields the accumulated sum of the elements in the slice
func AccumulateSlice[T intType](arr []T) iter.Seq[T] {
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

// Accumulate returns an iterator that yields the accumulated sum of the elements yielded by seq
func Accumulate[T intType](seq iter.Seq[T]) iter.Seq[T] {
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

// AccumulateFuncSlice returns an iterator that yields the accumulated result of applying f to the elements in the slice
func AccumulateFuncSlice[T any](arr []T, f func(T, T) T) iter.Seq[T] {
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

// AccumulateFunc returns an iterator that yields the accumulated result of applying f to the elements yielded by seq
func AccumulateFunc[T any](seq iter.Seq[T], f func(T, T) T) iter.Seq[T] {
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

// ChainSlice returns an iterator that yields the elements of the slice in order
func ChainSlice[T any](arr ...[]T) iter.Seq[T] {
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

// Chain returns an iterator that yields the elements yielded by seqs in order
func Chain[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
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

// DropSlice returns an iterator that yields elements not matching the predicate
func DropSlice[T any](arr []T, predicate func(T) bool) iter.Seq[T] {
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

// Drop returns an iterator that yields elements not matching the predicate
func Drop[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
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

// FilterSlice returns an iterator that yields elements matching the predicate
func FilterSlice[T any](arr []T, predicate func(T) bool) iter.Seq[T] {
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

// Filter returns an iterator that yields elements matching the predicate
func Filter[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
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

// PairwiseSlice returns an iterator that yields pairs of adjacent elements in the slice
// If the slice has less than 2 elements, the empty iterator is returned
func PairWiseSlice[T any](arr []T) iter.Seq[[2]T] {
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

// PairWise returns an iterator that yields pairs of adjacent elements yielded by seq
// If the sequence has less than 2 elements, the empty iterator is returned
func PairWise[T any](seq iter.Seq[T]) iter.Seq[[2]T] {
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

// WhileSlice returns an iterator that yields elements of the slice until the predicate is false
func WhileSlice[T any](arr []T, predicate func(T) bool) iter.Seq[T] {
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

// While returns an iterator that yields elements of the sequence until the predicate is false
func While[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
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

// LimitSlice returns an iterator that yields the up to the first n elements of the slice
func LimitSlice[T any](arr []T, n int) iter.Seq[T] {
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

// Limit returns an iterator that yields the up to the first n elements yielded by the sequence
func Limit[T any](seq iter.Seq[T], n int) iter.Seq[T] {
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

// ApplySlice returns an iterator that yields the result of applying f to each element in the slice
func ApplySlice[U, V any](arr []U, f func(U) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range arr {
			if !yield(f(v)) {
				break
			}
		}
	}
}

// Apply returns an iterator that yields the result of applying f to each element yielded by the sequence
func Apply[U, V any](seq iter.Seq[U], f func(U) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !yield(f(v)) {
				break
			}
		}
	}
}

// Unpack returns two iterators that yield the keys and values of the sequence
func Unpack[U, V any](seq iter.Seq2[U, V]) (iter.Seq[U], iter.Seq[V]) {
	keys := func(yield func(U) bool) {
		// nolint:revive
		for k, _ := range seq {
			if !yield(k) {
				break
			}
		}
	}
	vals := func(yield func(V) bool) {
		for _, v := range seq {
			if !yield(v) {
				break
			}
		}
	}
	return keys, vals
}

// UnpackMap returns two iterators that yield the keys and values of the map
func UnpackMap[U comparable, V any](in map[U]V) (iter.Seq[U], iter.Seq[V]) {
	keys := func(yield func(U) bool) {
		for k := range in {
			if !yield(k) {
				break
			}
		}
	}
	vals := func(yield func(V) bool) {
		for _, v := range in {
			if !yield(v) {
				break
			}
		}
	}
	return keys, vals
}

// Index returns an iterator that yields the index and value of the yielded elements from the sequence
func Index[T any](seq iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for v := range seq {
			if !yield(i, v) {
				break
			}
			i++
		}
	}
}

// RangeOver returns an iterator that yields the elements of the sequence from start to stop incrementing by step
func RangeOver[T any](seq iter.Seq[T], start, step, stop int) iter.Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		k := 0
		for v := range seq {
			if i >= stop {
				break
			}
			if i < start {
				i++
				continue
			}
			if k%step == 0 {
				if !yield(v) {
					break
				}
			}
			i++
			k++
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
