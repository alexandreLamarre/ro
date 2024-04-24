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

func TestCycle(t *testing.T) {
	s := []string{"a", "b", "c"}
	res := []string{}
	n := len(s) * 2
	for v := range ro.Cycle(s) {
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
	for v := range ro.Cycle(nan) {
		res2 = append(res2, v)
		n--
		if n == 0 {
			break
		}
	}
	assert.Equal(t, []int{}, res2)
}
