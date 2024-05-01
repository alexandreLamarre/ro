//go:build goexperiment.rangefunc

package ro

import (
	"iter"
	"math/bits"

	"github.com/samber/lo"
)

// Zip returns an iterator that yields the elements of arr1 and arr2 as a tuple
func Zip[U, V any](arr1 []U, arr2 []V) iter.Seq[lo.Tuple2[U, V]] {
	return func(yield func(lo.Tuple2[U, V]) bool) {
		for i := 0; i < len(arr1) && i < len(arr2); i++ {

			if !yield(lo.Tuple2[U, V]{A: arr1[i], B: arr2[i]}) {
				break
			}
		}
	}
}

// ZipIter returns an iterator that yields the elements of seq1 and seq2 as a tuple
func ZipIter[U, V any](seq1 iter.Seq[U], seq2 iter.Seq[V]) iter.Seq[lo.Tuple2[U, V]] {
	return func(yield func(lo.Tuple2[U, V]) bool) {
		p1, stop1 := iter.Pull(seq1)
		defer stop1()
		p2, stop2 := iter.Pull(seq2)
		defer stop2()
		for {
			var val lo.Tuple2[U, V]
			var ok1, ok2 bool
			val.A, ok1 = p1()
			val.B, ok2 = p2()
			if (!ok1 && !ok2) || !yield(val) {
				return
			}
		}
	}
}

// ZipFill returns an iterator that yields the elements of arr1 and arr2 as a tuple, padding missing values as necessary
func ZipFill[U, V any](arr1 []U, arr2 []V, fillU U, fillV V) iter.Seq[lo.Tuple2[U, V]] {
	return func(yield func(lo.Tuple2[U, V]) bool) {
		for i := 0; i < len(arr1) || i < len(arr2); i++ {
			var val lo.Tuple2[U, V]
			if i < len(arr1) {
				val.A = arr1[i]
			} else {
				val.A = fillU
			}
			if i < len(arr2) {
				val.B = arr2[i]
			} else {
				val.B = fillV
			}
			if !yield(val) {
				break
			}
		}
	}
}

// ZipFill returns an iterator that yields the elements of seq1 and seq2 as a tuple, padding missing values as necessary
func ZipFillIter[U, V any](seq1 iter.Seq[U], seq2 iter.Seq[V], fillU U, fillV V) iter.Seq[lo.Tuple2[U, V]] {
	return func(yield func(lo.Tuple2[U, V]) bool) {
		p1, stop1 := iter.Pull(seq1)
		defer stop1()
		p2, stop2 := iter.Pull(seq2)
		defer stop2()
		for {
			var val lo.Tuple2[U, V]
			var ok1, ok2 bool
			val.A, ok1 = p1()
			val.B, ok2 = p2()
			// TODO : we should also check that val.A/ val.B is indeed the zero element per the iter.Pull convention
			if !ok1 {
				val.A = fillU
			}
			if !ok2 {
				val.B = fillV
			}
			if !ok1 && !ok2 {
				return
			}
			if !yield(val) {
				return
			}
		}
	}
}

// Product returns an iterator that yields the cartesian product of slice1 and slice2 as a tuple.
// The returned iterator of tuples is ordered [A1, B1], [A1, B2], [A2, B1], [A2, B2], ...
func Product[U, V any](arr1 []U, arr2 []V) iter.Seq[lo.Tuple2[U, V]] {
	return func(yield func(lo.Tuple2[U, V]) bool) {
		for _, u := range arr1 {
			for _, v := range arr2 {
				if !yield(lo.Tuple2[U, V]{A: u, B: v}) {
					return
				}
			}
		}
	}
}

// ProductIter returns an iterator that yields the cartesian product of seq1 and seq2 as a tuple.
// The returned iterator of tuples is ordered [A1, B1], [A1, B2], [A2, B1], [A2, B2], ...
func ProductIter[U, V any](seq1 iter.Seq[U], seq2 iter.Seq[V]) iter.Seq[lo.Tuple2[U, V]] {
	return func(yield func(lo.Tuple2[U, V]) bool) {
		for u := range seq1 {
			for v := range seq2 {
				if !yield(lo.Tuple2[U, V]{A: u, B: v}) {
					return
				}
			}
		}
	}
}

// Permutations returns an iterator that yields all possible n-length permutations of the slice
// If k <= 0, the empty iterator is returned.
// If k > len(seq), the iterator yields all permutations of seq
//
// Note : allocates a state buffer of the same size as the input to keep track of visited permutations
// as we iterate through them.
func Permutations[T any](arr []T, k int) iter.Seq[[]T] {
	if k <= 0 {
		return empty[[]T]()
	}

	if k > len(arr) {
		k = len(arr)
	}

	return func(yield func([]T) bool) {
		n := len(arr)
		output := make([]T, k)
		copy(output, arr)
		if !yield(output) {
			return
		}
		state := make([]int, n)
		i := 0
		for i < n {
			if state[i] < i {
				if i%2 == 0 {
					arr[0], arr[i] = arr[i], arr[0]
				} else {
					arr[state[i]], arr[i] = arr[i], arr[state[i]]
				}
				output := make([]T, k)
				copy(output, arr)
				if !yield(output) {
					return
				}
				state[i]++
				i = 0
			} else {
				state[i] = 0
				i++
			}
		}
	}
}

// Combinations returns an iterator that yields all possible k-length combinations of the slice
// If k <= 0, the empty iterator is returned.
// If k > len(seq), the iterator yields all combinations of seq
func Combinations[T any](arr []T, k int) iter.Seq[[]T] {
	length := uint(len(arr))
	if k <= 0 {
		return empty[[]T]()
	}
	if k > len(arr) {
		k = len(arr)
	}
	return func(yield func([]T) bool) {
		for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
			if k > 0 && bits.OnesCount(uint(subsetBits)) != k {
				continue
			}

			var subset []T

			for object := uint(0); object < length; object++ {
				if (subsetBits>>object)&1 == 1 {
					subset = append(subset, arr[object])
				}
			}
			if !yield(subset) {
				break
			}
		}
	}
}
