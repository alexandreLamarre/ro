//go:build goexperiment.rangefunc

package ro

import (
	"iter"

	"github.com/samber/lo"
)

// Zip returns an iterator that yields the elements of seq1 and seq2 as a tuple
func Zip[U, V any](seq1 []U, seq2 []V) iter.Seq[lo.Tuple2[U, V]] {
	return func(yield func(lo.Tuple2[U, V]) bool) {
		for i := 0; i < len(seq1) && i < len(seq2); i++ {

			if !yield(lo.Tuple2[U, V]{A: seq1[i], B: seq2[i]}) {
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

// ZipFill returns an iterator that yields the elements of seq1 and seq2 as a tuple, padding missing values as necessary
func ZipFill[U, V any](seq1 []U, seq2 []V, fillU U, fillV V) iter.Seq[lo.Tuple2[U, V]] {
	return func(yield func(lo.Tuple2[U, V]) bool) {
		for i := 0; i < len(seq1) || i < len(seq2); i++ {
			var val lo.Tuple2[U, V]
			if i < len(seq1) {
				val.A = seq1[i]
			} else {
				val.A = fillU
			}
			if i < len(seq2) {
				val.B = seq2[i]
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

// Product returns an iterator that yields the cartesian product of seq1 and seq2 as a tuple.
// The returned iterator of tuples is ordered [A1, B1], [A1, B2], [A2, B1], [A2, B2], ...
func Product[U, V any](seq1 []U, seq2 []V) iter.Seq[lo.Tuple2[U, V]] {
	return func(yield func(lo.Tuple2[U, V]) bool) {
		for _, u := range seq1 {
			for _, v := range seq2 {
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

// Permutations returns an iterator that yields all possible n-length permutations of seq
// If k <= 0, the empty iterator is returned.
//
// Note : allocates a state buffer of the same size as the input to keep track of visited permutations
// as we iterate through them.
func Permutations[T any](seq []T, k int) iter.Seq[[]T] {
	if k <= 0 {
		return empty[[]T]()
	}

	return func(yield func([]T) bool) {
		n := len(seq)
		output := make([]T, k)
		copy(output, seq)
		if !yield(output) {
			return
		}
		state := make([]int, n)
		i := 0
		for i < n {
			if state[i] < i {
				if i%2 == 0 {
					seq[0], seq[i] = seq[i], seq[0]
				} else {
					seq[state[i]], seq[i] = seq[i], seq[state[i]]
				}
				output := make([]T, k)
				copy(output, seq)
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
