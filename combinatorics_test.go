package ro_test

import (
	"testing"

	"github.com/alexandreLamarre/ro"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestZipSlice(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.ZipSlice([]int{1, 2, 3}, []int{4, 5, 6}) {
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

	res2 := []lo.Tuple2[int, int]{}
	for v := range ro.ZipSlice([]int{1, 2, 3}, []int{4, 5, 6}) {
		res2 = append(res2, v)
		break
	}

	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 4,
		},
	}, res2)
}

func TestZip(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.Zip(ro.FromSlice([]int{1, 2, 3}), ro.FromSlice([]int{4, 5, 6})) {
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

	res2 := []lo.Tuple2[int, int]{}
	for v := range ro.Zip(ro.FromSlice([]int{1, 2, 3}), ro.FromSlice([]int{4, 5, 6})) {
		res2 = append(res2, v)
		break
	}
	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 4,
		},
	}, res2)
}

func TestZipFillSlice(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.ZipFillSlice([]int{1, 2, 3}, []int{4, 5}, -1, -2) {
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
	for v := range ro.ZipFillSlice([]int{1, 2}, []int{4, 5, 6}, -1, -2) {
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

	res3 := []lo.Tuple2[int, int]{}
	for v := range ro.ZipFillSlice([]int{1, 2}, []int{1, 2}, -1, -2) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 1,
		},
	}, res3)
}

func TestZipFill(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.ZipFill(ro.FromSlice([]int{1, 2, 3}), ro.FromSlice([]int{4, 5}), -1, -2) {
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
	for v := range ro.ZipFill(ro.FromSlice([]int{1, 2}), ro.FromSlice([]int{4, 5, 6}), -1, -2) {
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

	res3 := []lo.Tuple2[int, int]{}
	for v := range ro.ZipFill(ro.FromSlice([]int{1, 2}), ro.FromSlice([]int{1, 2}), -1, -2) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 1,
		},
	}, res3)
}

func TestProductSlice(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.ProductSlice([]int{1, 2, 3}, []int{4, 5}) {
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

	res2 := []lo.Tuple2[int, int]{}
	for v := range ro.ProductSlice([]int{1, 2, 3}, []int{4, 5}) {
		res2 = append(res2, v)
		break
	}
	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 4,
		},
	}, res2)
}

func TestProduct(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.Product(ro.FromSlice([]int{1, 2, 3}), ro.FromSlice([]int{4, 5})) {
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

	res2 := []lo.Tuple2[int, int]{}
	for v := range ro.Product(ro.FromSlice([]int{1, 2, 3}), ro.FromSlice([]int{4, 5})) {
		res2 = append(res2, v)
		break
	}

	assert.Equal(t, []lo.Tuple2[int, int]{
		{
			A: 1,
			B: 4,
		},
	}, res2)
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

	res5 := [][]int{}
	for v := range ro.Permutations([]int{1, 2, 3}, 3) {
		res5 = append(res5, v)
		break
	}
	assert.Equal(t, [][]int{{1, 2, 3}}, res5)

	res6 := [][]int{}
	for v := range ro.Permutations([]int{1, 2, 3}, 5) {
		res6 = append(res6, v)
	}
	assert.Equal(t, [][]int{{1, 2, 3}, {2, 1, 3}, {3, 1, 2}, {1, 3, 2}, {2, 3, 1}, {3, 2, 1}}, res6)
}

func TestCombinations(t *testing.T) {
	res := [][]string{}
	for v := range ro.Combinations([]string{"A", "B", "C", "D"}, 2) {
		res = append(res, v)
	}
	assert.Equal(t, [][]string{{"A", "B"}, {"A", "C"}, {"B", "C"}, {"A", "D"}, {"B", "D"}, {"C", "D"}}, res)

	res2 := [][]string{}
	for v := range ro.Combinations([]string{"A", "B", "C", "D"}, 0) {
		res2 = append(res2, v)
	}
	assert.Equal(t, [][]string{}, res2)

	res3 := [][]string{}
	for v := range ro.Combinations([]string{"A", "B", "C", "D"}, 2) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, [][]string{{"A", "B"}}, res3)

	res4 := [][]string{}
	for v := range ro.Combinations([]string{"A", "B", "C", "D"}, 5) {
		res4 = append(res4, v)
	}
	assert.Equal(t, [][]string{
		{
			"A",
			"B",
			"C",
			"D",
		},
	}, res4)
}
