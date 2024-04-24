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
