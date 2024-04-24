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
