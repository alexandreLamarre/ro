//go:build goexperiment.rangefunc

package ro_test

import (
	"testing"

	"github.com/alexandreLamarre/ro"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.Zip([]int{1, 2, 3}, []int{4, 5, 6}) {
		res = append(res, v)
	}
	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 4,
		},
		{
			A: 2,
			B: 5,
		},
		{
			A: 3,
			B: 6,
		},
	}, res)
}

func TestZipIter(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.ZipIter(ro.SeqAsIter([]int{1, 2, 3}), ro.SeqAsIter([]int{4, 5, 6})) {
		res = append(res, v)
	}
	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 4,
		},
		{
			A: 2,
			B: 5,
		},
		{
			A: 3,
			B: 6,
		},
	}, res)
}

func TestZipFill(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.ZipFill([]int{1, 2, 3}, []int{4, 5}, -1, -2) {
		res = append(res, v)
	}
	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 4,
		},
		{
			A: 2,
			B: 5,
		},
		{
			A: 3,
			B: -2,
		},
	}, res)

	res2 := []lo.Tuple2[int, int]{}
	for v := range ro.ZipFill([]int{1, 2}, []int{4, 5, 6}, -1, -2) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 4,
		},
		{
			A: 2,
			B: 5,
		},
		{
			A: -1,
			B: 6,
		},
	}, res2)
}

func TestZipFillIter(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.ZipFillIter(ro.SeqAsIter([]int{1, 2, 3}), ro.SeqAsIter([]int{4, 5}), -1, -2) {
		res = append(res, v)
	}
	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 4,
		},
		{
			A: 2,
			B: 5,
		},
		{
			A: 3,
			B: -2,
		},
	}, res)

	res2 := []lo.Tuple2[int, int]{}
	for v := range ro.ZipFillIter(ro.SeqAsIter([]int{1, 2}), ro.SeqAsIter([]int{4, 5, 6}), -1, -2) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 4,
		},
		{
			A: 2,
			B: 5,
		},
		{
			A: -1,
			B: 6,
		},
	}, res2)
}

func TestProduct(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.Product([]int{1, 2, 3}, []int{4, 5}) {
		res = append(res, v)
	}
	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 4,
		},
		{
			A: 1,
			B: 5,
		},
		{
			A: 2,
			B: 4,
		},
		{
			A: 2,
			B: 5,
		},
		{
			A: 3,
			B: 4,
		},
		{
			A: 3,
			B: 5,
		},
	}, res)
}

func TestProductIter(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.ProductIter(ro.SeqAsIter([]int{1, 2, 3}), ro.SeqAsIter([]int{4, 5})) {
		res = append(res, v)
	}
	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 4,
		},
		{
			A: 1,
			B: 5,
		},
		{
			A: 2,
			B: 4,
		},
		{
			A: 2,
			B: 5,
		},
		{
			A: 3,
			B: 4,
		},
		{
			A: 3,
			B: 5,
		},
	}, res)
}

func TestPermutations(t *testing.T) {
	res := [][]int{}
	for v := range ro.Permutations([]int{1, 2, 3}, 2) {
		res = append(res, v)
	}
	assert.Equal(t, [][]int{{1, 2}, {2, 1}, {3, 1}, {1, 3}, {2, 3}, {3, 2}}, res)

	res2 := [][]int{}
	for v := range ro.Permutations([]int{1, 2, 3}, 3) {
		res2 = append(res2, v)
	}
	assert.Equal(t, [][]int{{1, 2, 3}, {2, 1, 3}, {3, 1, 2}, {1, 3, 2}, {2, 3, 1}, {3, 2, 1}}, res2)

	res3 := [][]int{}
	for v := range ro.Permutations([]int{1, 2, 3}, 0) {
		res3 = append(res3, v)
	}
	assert.Equal(t, [][]int{}, res3)

	res4 := [][]int{}
	for v := range ro.Permutations([]int{1, 2, 3}, 3) {
		res4 = append(res4, v)
	}
	assert.Equal(t, [][]int{{1, 2, 3}, {2, 1, 3}, {3, 1, 2}, {1, 3, 2}, {2, 3, 1}, {3, 2, 1}}, res4)
}
