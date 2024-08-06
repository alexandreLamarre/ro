package ro_test

import (
	"testing"

	"github.com/alexandreLamarre/ro"
	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	res := []int{}
	n := 3
	for v := range ro.Count(1, 4) {
		res = append(res, v)
		n--
		if n == 0 {
			break
		}
	}
	assert.Equal(t, []int{1, 5, 9}, res)

	res2 := []int{}
	n = 5
	for v := range ro.Count(0, 1) {
		res2 = append(res2, v)
		n--
		if n == 0 {
			break
		}
	}
	assert.Equal(t, []int{0, 1, 2, 3, 4}, res2)

	res3 := []int8{}
	n = 3
	for v := range ro.Count(int8(1), int8(1)) {
		res3 = append(res3, v)
		n--
		if n == 0 {
			break
		}
	}
	assert.Equal(t, []int8{1, 2, 3}, res3)
}

func TestCycleSlice(t *testing.T) {
	s := []string{"a", "b", "c"}
	res := []string{}
	n := len(s) * 2
	for v := range ro.CycleSlice(s) {
		res = append(res, v)
		n--
		if n == 0 {
			break
		}
	}
	assert.Equal(t, []string{"a", "b", "c", "a", "b", "c"}, res)

	nan := []int{}
	res2 := []int{}
	n = 10
	for v := range ro.CycleSlice(nan) {
		res2 = append(res2, v)
		n--
		if n == 0 {
			break
		}
	}
	assert.Equal(t, []int{}, res2)
}

func TestCycle(t *testing.T) {
	s := []string{"a", "b", "c"}
	res := []string{}
	n := len(s) * 2
	for v := range ro.Cycle(ro.FromSlice(s)) {
		res = append(res, v)
		n--
		if n == 0 {
			break
		}
	}
	assert.Equal(t, []string{"a", "b", "c", "a", "b", "c"}, res)

	// TODO : requires a redesign of the implementation of CycleIter to prevent infinite loops
	// nan := []int{}
	// res2 := []int{}
	// n = 10
	// for v := range ro.CycleIter(ro.SeqAsIter(nan)) {
	// 	res2 = append(res2, v)
	// 	n--
	// 	if n == 0 {
	// 		break
	// 	}
	// }
	// assert.Equal(t, []int{}, res2)
}

func TestRepeat(t *testing.T) {
	res := []int{}
	n := 3
	for v := range ro.Repeat(1) {
		res = append(res, v)
		n--
		if n == 0 {
			break
		}
	}
	assert.Equal(t, []int{1, 1, 1}, res)

	res2 := []int8{}
	n = 5
	for v := range ro.Repeat(int8(2)) {
		res2 = append(res2, v)
		n--
		if n == 0 {
			break
		}
	}
	assert.Equal(t, []int8{2, 2, 2, 2, 2}, res2)
}
