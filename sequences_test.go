//go:build goexperiment.rangefunc

package ro_test

import (
	"fmt"
	"iter"
	"sync"
	"testing"

	"github.com/alexandreLamarre/ro"
	"github.com/stretchr/testify/assert"
)

func TestAccumulateSlice(t *testing.T) {
	seq := []int{1, 2, 3, 4}
	res := []int{}
	for v := range ro.AccumulateSlice(seq) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 3, 6, 10}, res)

	empty := []int{}
	res2 := []int{}
	for v := range ro.AccumulateSlice(empty) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	seq3 := []int{1, 2, 3, 4}
	res3 := []int{}
	for v := range ro.AccumulateSlice(seq3) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []int{1}, res3)
}

func TestAccumulate(t *testing.T) {
	res := []int{}
	for v := range ro.Accumulate(ro.FromSlice([]int{1, 2, 3, 4})) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 3, 6, 10}, res)

	res2 := []int{}
	for v := range ro.Accumulate(ro.FromSlice([]int{})) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	res3 := []int{}
	for v := range ro.Accumulate(ro.FromSlice([]int{1, 2, 3, 4})) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []int{1}, res3)
}

func TestAccumulateFuncSlice(t *testing.T) {
	res := [][]string{}
	for v := range ro.AccumulateFuncSlice([][]string{{"a", "b"}, {"c", "d"}}, func(a, b []string) []string { return append(a, b...) }) {
		res = append(res, v)
	}
	assert.Equal(t, [][]string{{"a", "b"}, {"a", "b", "c", "d"}}, res)

	res2 := []int{}
	for v := range ro.AccumulateFuncSlice([]int{}, func(a, b int) int { return a % b }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	res3 := [][]string{}
	for v := range ro.AccumulateFuncSlice([][]string{{"a", "b"}, {"c", "d"}}, func(a, b []string) []string { return append(a, b...) }) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, [][]string{{"a", "b"}}, res3)
}

func TestAccumulateFunc(t *testing.T) {
	res := [][]string{}
	for v := range ro.AccumulateFunc(ro.FromSlice([][]string{{"a", "b"}, {"c", "d"}}), func(a, b []string) []string { return append(a, b...) }) {
		res = append(res, v)
	}
	assert.Equal(t, [][]string{{"a", "b"}, {"a", "b", "c", "d"}}, res)

	res2 := []int{}
	for v := range ro.AccumulateFunc(ro.FromSlice([]int{}), func(a, b int) int { return a % b }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	res3 := [][]string{}
	for v := range ro.AccumulateFunc(ro.FromSlice([][]string{{"a", "b"}, {"c", "d"}}), func(a, b []string) []string { return append(a, b...) }) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, [][]string{{"a", "b"}}, res3)
}

func TestBatchSlice(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := [][]int{}
	for v := range ro.BatchSlice(seq, 3) {
		res = append(res, v)
	}
	assert.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, res)

	empty := []int{}
	res2 := [][]int{}
	for v := range ro.BatchSlice(empty, 3) {
		res2 = append(res2, v)
	}
	assert.Equal(t, [][]int{}, res2)

	res3 := [][]int{}
	for v := range ro.BatchSlice(seq, 3) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, [][]int{{1, 2, 3}}, res3)
}

func TestBatch(t *testing.T) {
	res := [][]int{}
	for v := range ro.Batch(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), 3) {
		res = append(res, v)
	}
	assert.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, res)

	res2 := [][]int{}
	for v := range ro.Batch(ro.FromSlice([]int{}), 3) {
		res2 = append(res2, v)
	}
	assert.Equal(t, [][]int{}, res2)

	res3 := [][]int{}
	for v := range ro.Batch(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), 3) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, [][]int{{1, 2, 3}}, res3)

	res4 := [][]int{}
	for v := range ro.Batch(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), 3) {
		res4 = append(res4, v)
	}
	assert.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}, res4)
}

func TestChainSlice(t *testing.T) {
	res := []int{}
	for v := range ro.ChainSlice([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, res)

	res2 := []int{}
	for v := range ro.ChainSlice([]int{}, []int{}, []int{}) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	res3 := []int{}
	for v := range ro.ChainSlice([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []int{1}, res3)
}

func TestChain(t *testing.T) {
	res := []int{}
	for v := range ro.Chain(ro.FromSlice([]int{1, 2, 3}), ro.FromSlice([]int{4, 5, 6}), ro.FromSlice([]int{7, 8, 9})) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, res)

	res2 := []int{}
	for v := range ro.Chain(ro.FromSlice([]int{}), ro.FromSlice([]int{}), ro.FromSlice([]int{})) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	res3 := []int{}
	for v := range ro.Chain(ro.FromSlice([]int{1, 2, 3}), ro.FromSlice([]int{4, 5, 6}), ro.FromSlice([]int{7, 8, 9})) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []int{1}, res3)
}

func TestDropSlice(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := []int{}
	for v := range ro.DropSlice(seq, func(i int) bool { return i%3 == 0 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 4, 5, 7, 8}, res)

	empty := []int{}
	res2 := []int{}
	for v := range ro.DropSlice(empty, func(i int) bool { return i%3 == 0 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	res3 := []int{}
	for v := range ro.DropSlice(seq, func(i int) bool { return i%3 == 0 }) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []int{1}, res3)
}

func TestDrop(t *testing.T) {
	res := []int{}
	for v := range ro.Drop(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), func(i int) bool { return i%3 == 0 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 4, 5, 7, 8}, res)

	res2 := []int{}
	for v := range ro.Drop(ro.FromSlice([]int{}), func(i int) bool { return i%3 == 0 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	res3 := []int{}
	for v := range ro.Drop(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), func(i int) bool { return i%3 == 0 }) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []int{1}, res3)
}

func TestFilterSlice(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := []int{}
	for v := range ro.FilterSlice(seq, func(i int) bool { return i%3 == 0 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{3, 6, 9}, res)

	empty := []int{}
	res2 := []int{}
	for v := range ro.FilterSlice(empty, func(i int) bool { return i%3 == 0 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	res3 := []int{}
	for v := range ro.FilterSlice(seq, func(i int) bool { return i%3 == 0 }) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []int{3}, res3)
}

func TestFilter(t *testing.T) {
	res := []int{}
	for v := range ro.Filter(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), func(i int) bool { return i%3 == 0 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{3, 6, 9}, res)

	res2 := []int{}
	for v := range ro.Filter(ro.FromSlice([]int{}), func(i int) bool { return i%3 == 0 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	res3 := []int{}
	for v := range ro.Filter(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), func(i int) bool { return i%3 == 0 }) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []int{3}, res3)
}

func TestPairwiseSlice(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := [][2]int{}
	for v := range ro.PairWiseSlice(seq) {
		res = append(res, v)
	}
	assert.Equal(t, [][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}}, res)

	empty := []int{}
	res2 := [][2]int{}
	for v := range ro.PairWiseSlice(empty) {
		res2 = append(res2, v)
	}
	assert.Equal(t, [][2]int{}, res2)

	one := []int{1}
	res3 := [][2]int{}
	for v := range ro.PairWiseSlice(one) {
		res3 = append(res3, v)
	}
	assert.Equal(t, [][2]int{}, res3)

	res4 := [][2]int{}
	for v := range ro.PairWiseSlice(seq) {
		res4 = append(res4, v)
		break
	}
	assert.Equal(t, [][2]int{{1, 2}}, res4)
}

func TestPairwise(t *testing.T) {
	res := [][2]int{}
	for v := range ro.PairWise(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})) {
		res = append(res, v)
	}
	assert.Equal(t, [][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}}, res)

	res2 := [][2]int{}
	for v := range ro.PairWise(ro.FromSlice([]int{})) {
		res2 = append(res2, v)
	}
	assert.Equal(t, [][2]int{}, res2)

	one := []int{1}
	res3 := [][2]int{}
	for v := range ro.PairWise(ro.FromSlice(one)) {
		res3 = append(res3, v)
	}
	assert.Equal(t, [][2]int{}, res3)

	res4 := [][2]int{}
	for v := range ro.PairWise(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})) {
		res4 = append(res4, v)
		break
	}
	assert.Equal(t, [][2]int{{1, 2}}, res4)
}

func TestWhileSlice(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := []int{}
	for v := range ro.WhileSlice(seq, func(i int) bool { return i < 5 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4}, res)

	empty := []int{}
	res2 := []int{}
	for v := range ro.WhileSlice(empty, func(i int) bool { return i < 5 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	all := []int{1, 2, 3, 4, 5}
	res3 := []int{}
	for v := range ro.WhileSlice(all, func(i int) bool { return i < 6 }) {
		res3 = append(res3, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res3)

	res4 := []int{}
	for v := range ro.WhileSlice(seq, func(i int) bool { return i < 5 }) {
		res4 = append(res4, v)
		break
	}
	assert.Equal(t, []int{1}, res4)
}

func TestWhile(t *testing.T) {
	res := []int{}
	for v := range ro.While(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), func(i int) bool { return i < 5 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4}, res)

	res2 := []int{}
	for v := range ro.While(ro.FromSlice([]int{}), func(i int) bool { return i < 5 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	all := []int{1, 2, 3, 4, 5}
	res3 := []int{}
	for v := range ro.While(ro.FromSlice(all), func(i int) bool { return i < 6 }) {
		res3 = append(res3, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res3)

	res4 := []int{}
	for v := range ro.While(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), func(i int) bool { return i < 5 }) {
		res4 = append(res4, v)
		break
	}
	assert.Equal(t, []int{1}, res4)
}

func TestLimitSlice(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := []int{}
	for v := range ro.LimitSlice(seq, 5) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res)

	empty := []int{}
	res2 := []int{}
	for v := range ro.LimitSlice(empty, 5) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	all := []int{1, 2, 3, 4, 5}
	res3 := []int{}
	for v := range ro.LimitSlice(all, 10) {
		res3 = append(res3, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res3)

	res4 := []int{}
	for v := range ro.LimitSlice(seq, 5) {
		res4 = append(res4, v)
		break
	}
	assert.Equal(t, []int{1}, res4)
}

func TestLimit(t *testing.T) {
	res := []int{}
	for v := range ro.Limit(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), 5) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res)
	res2 := []int{}
	for v := range ro.Limit(ro.FromSlice([]int{}), 5) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	all := []int{1, 2, 3, 4, 5}
	res3 := []int{}
	for v := range ro.Limit(ro.FromSlice(all), 10) {
		res3 = append(res3, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res3)

	res4 := []int{}
	for v := range ro.Limit(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), 5) {
		res4 = append(res4, v)
		break
	}
	assert.Equal(t, []int{1}, res4)
}

func TestApplySlice(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5}
	res := []int{}
	for v := range ro.ApplySlice(seq, func(i int) int { return i * 2 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{2, 4, 6, 8, 10}, res)

	empty := []int{}
	res2 := []int{}
	for v := range ro.ApplySlice(empty, func(i int) int { return i * 2 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	res3 := []int{}
	for v := range ro.ApplySlice(seq, func(i int) int { return i * 2 }) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []int{2}, res3)
}

func TestApply(t *testing.T) {
	res := []int{}
	for v := range ro.Apply(ro.FromSlice([]int{1, 2, 3, 4, 5}), func(i int) int { return i * 2 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{2, 4, 6, 8, 10}, res)

	res2 := []int{}
	for v := range ro.Apply(ro.FromSlice([]int{}), func(i int) int { return i * 2 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	res3 := []int{}
	for v := range ro.Apply(ro.FromSlice([]int{1, 2, 3, 4, 5}), func(i int) int { return i * 2 }) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []int{2}, res3)
}

func sliceFullIter[T any](arr []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range arr {
			if !yield(i, v) {
				break
			}
		}

	}
}

func TestUnpack(t *testing.T) {
	resK := []int{}
	resV := []int{}
	ik, iv := ro.Unpack(sliceFullIter([]int{6, 5, 4, 3, 2, 1}))
	for v := range ik {
		resK = append(resK, v)
	}
	for v := range iv {
		resV = append(resV, v)
	}
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, resK)
	assert.Equal(t, []int{6, 5, 4, 3, 2, 1}, resV)

	// if values are interrupted before keys, the keys should be unaffected

	resK2 := []int{}
	resV2 := []int{}
	ik2, iv2 := ro.Unpack(sliceFullIter([]int{6, 5, 4, 3, 2, 1}))

	for v := range iv2 {
		resV2 = append(resV2, v)
		break
	}

	for k := range ik2 {
		resK2 = append(resK2, k)
	}
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, resK2)
	assert.Equal(t, []int{6}, resV2)

	// if keys are interrupted before values, the values should be unaffected

	resK3 := []int{}
	resV3 := []int{}
	ik3, iv3 := ro.Unpack(sliceFullIter([]int{6, 5, 4, 3, 2, 1}))

	for k := range ik3 {
		resK3 = append(resK3, k)
		break
	}

	for v := range iv3 {
		resV3 = append(resV3, v)
	}

	assert.Equal(t, []int{0}, resK3)
	assert.Equal(t, []int{6, 5, 4, 3, 2, 1}, resV3)
}

func TestUnpackMap(t *testing.T) {
	resK := []int{}
	resV := []int{}
	ik, iv := ro.UnpackMap(map[int]int{1: 6, 2: 5, 3: 4, 4: 3, 5: 2, 6: 1})
	for v := range ik {
		resK = append(resK, v)
	}
	for v := range iv {
		resV = append(resV, v)
	}
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6}, resK)
	assert.ElementsMatch(t, []int{6, 5, 4, 3, 2, 1}, resV)

	// if values are interrupted before keys, the keys should be unaffected

	resK2 := []int{}
	resV2 := []int{}
	ik2, iv2 := ro.UnpackMap(map[int]int{1: 6, 2: 5, 3: 4, 4: 3, 5: 2, 6: 1})

	for v := range iv2 {
		resV2 = append(resV2, v)
		break
	}

	for k := range ik2 {
		resK2 = append(resK2, k)
	}
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6}, resK2)
	assert.Len(t, resV2, 1)

	// if keys are interrupted before values, the values should be unaffected

	resK3 := []int{}
	resV3 := []int{}
	ik3, iv3 := ro.UnpackMap(map[int]int{1: 6, 2: 5, 3: 4, 4: 3, 5: 2, 6: 1})

	for k := range ik3 {
		resK3 = append(resK3, k)
		break
	}

	for v := range iv3 {
		resV3 = append(resV3, v)
	}

	assert.Len(t, resK3, 1)
	assert.ElementsMatch(t, []int{6, 5, 4, 3, 2, 1}, resV3)

	resK4 := []int{}
	resV4 := []int{}

	ik4, iv4 := ro.UnpackMap(map[int]int{})
	for v := range ik4 {
		resK4 = append(resK4, v)
	}
	for v := range iv4 {
		resV4 = append(resV4, v)
	}
	assert.Equal(t, []int{}, resK4)
	assert.Equal(t, []int{}, resV4)
}

func TestRangeOver(t *testing.T) {
	res := []int{}
	it := ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
	for v := range ro.RangeOver(it, 1, 2, 10) {
		res = append(res, v)
	}
	assert.Equal(t, []int{2, 4, 6, 8}, res)

	res2 := []int{}
	for v := range ro.RangeOver(ro.FromSlice([]int{}), 0, 2, 5) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	res3 := []int{}
	for v := range ro.RangeOver(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), 0, 2, 10) {
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []int{1}, res3)
}

func TestIndex(t *testing.T) {
	res := []int{}
	ind := []int{}
	for i, v := range ro.Index(ro.FromSlice([]int{1, 2, 3, 4, 5})) {
		ind = append(ind, i)
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, ind)

	res2 := []int{}
	ind2 := []int{}
	for i, v := range ro.Index(ro.FromSlice([]int{})) {
		res2 = append(res2, v)
		ind2 = append(ind2, i)
	}
	assert.Equal(t, []int{}, res2)
	assert.Equal(t, []int{}, ind2)

	res3 := []int{}
	ind3 := []int{}
	for i, v := range ro.Index(ro.FromSlice([]int{1, 2, 3, 4, 5})) {
		ind3 = append(ind3, i)
		res3 = append(res3, v)
		break
	}
	assert.Equal(t, []int{1}, res3)
	assert.Equal(t, []int{0}, ind3)
}

func TestTee(t *testing.T) {
	iters := ro.Tee(ro.FromSlice([]int{1, 2, 3, 4, 5}), 3)

	res := []int{}
	res1 := []int{}
	res2 := []int{}
	for v := range iters[0] {
		res = append(res, v)
	}
	for v := range iters[1] {
		res1 = append(res1, v)
	}
	for v := range iters[2] {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res1)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res2)

	noiters := ro.Tee(ro.FromSlice([]int{1, 2, 3, 4}), -5)
	assert.Len(t, noiters, 0)

	type testStruct struct {
		val string
	}

	iters2 := ro.Tee(ro.FromSlice([]*testStruct{{val: "a"}, {val: "b"}, {val: "c"}}), 3)
	assert.Len(t, iters2, 3)
	resMut := [][]*testStruct{{}, {}, {}}

	var wg sync.WaitGroup
	for i, iterS := range iters2 {
		i := i
		iterS := iterS
		wg.Add(1)
		go func() {
			defer wg.Done()
			// nolint:ineffassign
			for item := range iterS {
				item = &testStruct{val: fmt.Sprintf("replace%d", i)}
				resMut[i] = append(resMut[i], item)
			}
		}()

	}
	wg.Wait()

	for i, v := range resMut {
		assert.Len(t, v, 3)
		for _, item := range v {
			assert.Equal(t, fmt.Sprintf("replace%d", i), item.val)
		}
	}
	iters3 := ro.FromSlice([]int{1, 2, 3, 4, 5})
	out := ro.Tee(iters3, 1)
	assert.Len(t, out, 1)
	resOut := []int{}
	for v := range out[0] {
		resOut = append(resOut, v)
		break
	}
	assert.Equal(t, []int{1}, resOut)
}
