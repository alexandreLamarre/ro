package ro_test

import (
	"iter"
	"testing"

	"github.com/alexandreLamarre/ro"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func empty[T any]() iter.Seq[T] {
	return func(_ func(T) bool) {}
}

func TestToSlice(t *testing.T) {
	rangeIter := ro.Range(0, 10, 1)
	res := ro.ToSlice(
		rangeIter,
	)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, res)

	emptyIter := empty[int]()
	res2 := ro.ToSlice(
		emptyIter,
	)
	assert.Equal(t, []int{}, res2)
}

func TestFromSlice(t *testing.T) {
	i := ro.FromSlice([]string{"a", "b", "c"})
	res := []string{}
	for v := range i {
		res = append(res, v)
	}
	assert.Equal(t, []string{"a", "b", "c"}, res)

	i2 := ro.FromSlice([]int{})

	res2 := []int{}
	for v := range i2 {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
}

func TestFromString(t *testing.T) {
	i := ro.FromString("abc")
	res := []rune{}
	for v := range i {
		res = append(res, v)
	}
	assert.Equal(t, []rune{'a', 'b', 'c'}, res)

	i2 := ro.FromString("")
	res2 := []rune{}
	for v := range i2 {
		res2 = append(res2, v)
	}
	assert.Equal(t, []rune{}, res2)

	i3 := ro.FromString("abc")
	res3 := []rune{}
	for v := range i3 {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []rune{'a'}, res3)
}

func iter2ToTuple[U, V any](seq iter.Seq2[U, V]) []lo.Tuple2[U, V] {
	res := []lo.Tuple2[U, V]{}
	for k, v := range seq {
		res = append(res, lo.Tuple2[U, V]{A: k, B: v})
	}
	return res
}

func TestExtend(t *testing.T) {
	i := ro.FromSlice([]int{1, 2, 3})
	kv := ro.Extend(i)

	assert.Equal(t, []lo.Tuple2[struct{}, int]{
		{A: struct{}{}, B: 1},
		{A: struct{}{}, B: 2},
		{A: struct{}{}, B: 3},
	}, iter2ToTuple(kv))

	i2 := ro.FromSlice([]string{})
	kv2 := ro.Extend(i2)
	assert.Equal(t, []lo.Tuple2[struct{}, string]{}, iter2ToTuple(kv2))

	i3 := ro.FromSlice([]int{1, 2, 3})
	kv3 := ro.Extend(i3)
	res3 := []lo.Tuple2[struct{}, int]{}
	for k, v := range kv3 {
		res3 = append(res3, lo.Tuple2[struct{}, int]{A: k, B: v})
		break
	}
	assert.Equal(t, []lo.Tuple2[struct{}, int]{
		{A: struct{}{}, B: 1},
	}, res3)
}
