//go:build goexperiment.rangefunc

package ro_test

import (
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
	for v := range ro.AccumulateIter(ro.SeqAsIter([]int{1, 2, 3, 4})) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 3, 6, 10}, res)

	res2 := []int{}
	for v := range ro.AccumulateIter(ro.SeqAsIter([]int{})) {
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
	for v := range ro.AccumulateIterFunc(ro.SeqAsIter([][]string{{"a", "b"}, {"c", "d"}}), func(a, b []string) []string { return append(a, b...) }) {
		res = append(res, v)
	}
	assert.Equal(t, [][]string{{"a", "b"}, {"a", "b", "c", "d"}}, res)

	res2 := []int{}
	for v := range ro.AccumulateIterFunc(ro.SeqAsIter([]int{}), func(a, b int) int { return a % b }) {
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
	for v := range ro.BatchIter(ro.SeqAsIter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), 3) {
		res = append(res, v)
	}
	assert.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, res)

	res2 := [][]int{}
	for v := range ro.BatchIter(ro.SeqAsIter([]int{}), 3) {
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
	for v := range ro.ChainIter(ro.SeqAsIter([]int{1, 2, 3}), ro.SeqAsIter([]int{4, 5, 6}), ro.SeqAsIter([]int{7, 8, 9})) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, res)

	res2 := []int{}
	for v := range ro.ChainIter(ro.SeqAsIter([]int{}), ro.SeqAsIter([]int{}), ro.SeqAsIter([]int{})) {
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
	for v := range ro.DropIter(ro.SeqAsIter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), func(i int) bool { return i%3 == 0 }) {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2, 4, 5, 7, 8}, res)

	res2 := []int{}
	for v := range ro.DropIter(ro.SeqAsIter([]int{}), func(i int) bool { return i%3 == 0 }) {
		res2 = append(res2, v)
	}
	assert.Equal(t, []int{}, res2)
}
