//go:build goexperiment.rangefunc

package ro_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alexandreLamarre/ro"
	"github.com/stretchr/testify/assert"
)

func TestAccumulate(t *testing.T) {
	seq := []int{1, 2, 3, 4}
	res := []int{}
	for v := range ro.Accumulate(seq) {
		res = append(res, v)
	}

	assert.Equal(t, []int{1, 3, 6, 10}, res)

	empty := []int{}
	res2 := []int{}
	for v := range ro.Accumulate(empty) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
}

func TestAccumulateIter(t *testing.T) {
	res := []int{}
	for v := range ro.AccumulateIter(ro.FromSlice([]int{1, 2, 3, 4})) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 3, 6, 10}, res)

	res2 := []int{}
	for v := range ro.AccumulateIter(ro.FromSlice([]int{})) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
}

func TestAccumulateFunc(t *testing.T) {
	res := [][]string{}
	for v := range ro.AccumulateFunc([][]string{{"a", "b"}, {"c", "d"}}, func(a, b []string) []string { return append(a, b...) }) {
		res = append(res, v)
	}
	assert.Equal(t, [][]string{{"a", "b"}, {"a", "b", "c", "d"}}, res)

	res2 := []int{}
	for v := range ro.AccumulateFunc([]int{}, func(a, b int) int { return a % b }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
}

func TestAccumulateIterFunc(t *testing.T) {
	res := [][]string{}
	for v := range ro.AccumulateIterFunc(ro.FromSlice([][]string{{"a", "b"}, {"c", "d"}}), func(a, b []string) []string { return append(a, b...) }) {
		res = append(res, v)
	}
	assert.Equal(t, [][]string{{"a", "b"}, {"a", "b", "c", "d"}}, res)

	res2 := []int{}
	for v := range ro.AccumulateIterFunc(ro.FromSlice([]int{}), func(a, b int) int { return a % b }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
}

func TestBatch(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := [][]int{}
	for v := range ro.Batch(seq, 3) {
		res = append(res, v)
	}
	assert.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, res)

	empty := []int{}
	res2 := [][]int{}
	for v := range ro.Batch(empty, 3) {
		res2 = append(res2, v)
	}
	assert.Equal(t, [][]int{}, res2)
}

func TestBatchIter(t *testing.T) {
	res := [][]int{}
	for v := range ro.BatchIter(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), 3) {
		res = append(res, v)
	}
	assert.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, res)

	res2 := [][]int{}
	for v := range ro.BatchIter(ro.FromSlice([]int{}), 3) {
		res2 = append(res2, v)
	}
	assert.Equal(t, [][]int{}, res2)
}

func TestChain(t *testing.T) {
	res := []int{}
	for v := range ro.Chain([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, res)

	res2 := []int{}
	for v := range ro.Chain([]int{}, []int{}, []int{}) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
}

func TestChainIter(t *testing.T) {
	res := []int{}
	for v := range ro.ChainIter(ro.FromSlice([]int{1, 2, 3}), ro.FromSlice([]int{4, 5, 6}), ro.FromSlice([]int{7, 8, 9})) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, res)

	res2 := []int{}
	for v := range ro.ChainIter(ro.FromSlice([]int{}), ro.FromSlice([]int{}), ro.FromSlice([]int{})) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
}

func TestDrop(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := []int{}
	for v := range ro.Drop(seq, func(i int) bool { return i%3 == 0 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 4, 5, 7, 8}, res)

	empty := []int{}
	res2 := []int{}
	for v := range ro.Drop(empty, func(i int) bool { return i%3 == 0 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
}

func TestDropIter(t *testing.T) {
	res := []int{}
	for v := range ro.DropIter(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), func(i int) bool { return i%3 == 0 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 4, 5, 7, 8}, res)

	res2 := []int{}
	for v := range ro.DropIter(ro.FromSlice([]int{}), func(i int) bool { return i%3 == 0 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
}

func TestFilter(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := []int{}
	for v := range ro.Filter(seq, func(i int) bool { return i%3 == 0 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{3, 6, 9}, res)

	empty := []int{}
	res2 := []int{}
	for v := range ro.Filter(empty, func(i int) bool { return i%3 == 0 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
}

func TestFilterI(t *testing.T) {
	res := []int{}
	for v := range ro.FilterIter(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), func(i int) bool { return i%3 == 0 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{3, 6, 9}, res)

	res2 := []int{}
	for v := range ro.FilterIter(ro.FromSlice([]int{}), func(i int) bool { return i%3 == 0 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
}

func TestPairwise(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := [][2]int{}
	for v := range ro.PairWise(seq) {
		res = append(res, v)
	}
	assert.Equal(t, [][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}}, res)

	empty := []int{}
	res2 := [][2]int{}
	for v := range ro.PairWise(empty) {
		res2 = append(res2, v)
	}
	assert.Equal(t, [][2]int{}, res2)

	one := []int{1}
	res3 := [][2]int{}
	for v := range ro.PairWise(one) {
		res3 = append(res3, v)
	}
	assert.Equal(t, [][2]int{}, res3)

}

func TestPairwiseIter(t *testing.T) {
	res := [][2]int{}
	for v := range ro.PairWiseIter(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})) {
		res = append(res, v)
	}
	assert.Equal(t, [][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}}, res)

	res2 := [][2]int{}
	for v := range ro.PairWiseIter(ro.FromSlice([]int{})) {
		res2 = append(res2, v)
	}
	assert.Equal(t, [][2]int{}, res2)

	one := []int{1}
	res3 := [][2]int{}
	for v := range ro.PairWiseIter(ro.FromSlice(one)) {
		res3 = append(res3, v)
	}
	assert.Equal(t, [][2]int{}, res3)
}

func TestWhile(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := []int{}
	for v := range ro.While(seq, func(i int) bool { return i < 5 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4}, res)

	empty := []int{}
	res2 := []int{}
	for v := range ro.While(empty, func(i int) bool { return i < 5 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	all := []int{1, 2, 3, 4, 5}
	res3 := []int{}
	for v := range ro.While(all, func(i int) bool { return i < 6 }) {
		res3 = append(res3, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res3)
}

func TestWhileIter(t *testing.T) {
	res := []int{}
	for v := range ro.WhileIter(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), func(i int) bool { return i < 5 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4}, res)

	res2 := []int{}
	for v := range ro.WhileIter(ro.FromSlice([]int{}), func(i int) bool { return i < 5 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	all := []int{1, 2, 3, 4, 5}
	res3 := []int{}
	for v := range ro.WhileIter(ro.FromSlice(all), func(i int) bool { return i < 6 }) {
		res3 = append(res3, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res3)
}

func TestLimit(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := []int{}
	for v := range ro.Limit(seq, 5) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res)

	empty := []int{}
	res2 := []int{}
	for v := range ro.Limit(empty, 5) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	all := []int{1, 2, 3, 4, 5}
	res3 := []int{}
	for v := range ro.Limit(all, 10) {
		res3 = append(res3, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res3)
}

func TestLimitIter(t *testing.T) {
	res := []int{}
	for v := range ro.LimitIter(ro.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), 5) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res)
	res2 := []int{}
	for v := range ro.LimitIter(ro.FromSlice([]int{}), 5) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)

	all := []int{1, 2, 3, 4, 5}
	res3 := []int{}
	for v := range ro.LimitIter(ro.FromSlice(all), 10) {
		res3 = append(res3, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res3)
}

func TestApply(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5}
	res := []int{}
	for v := range ro.Apply(seq, func(i int) int { return i * 2 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{2, 4, 6, 8, 10}, res)

	empty := []int{}
	res2 := []int{}
	for v := range ro.Apply(empty, func(i int) int { return i * 2 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
}

func TestApplyIter(t *testing.T) {
	res := []int{}
	for v := range ro.ApplyIter(ro.FromSlice([]int{1, 2, 3, 4, 5}), func(i int) int { return i * 2 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{2, 4, 6, 8, 10}, res)

	res2 := []int{}
	for v := range ro.ApplyIter(ro.FromSlice([]int{}), func(i int) int { return i * 2 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
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

}
