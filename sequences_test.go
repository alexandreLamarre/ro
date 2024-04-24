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
