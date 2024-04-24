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
}
