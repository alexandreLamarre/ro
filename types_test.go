//go:build goexperiment.rangefunc

package ro_test

import (
	"testing"

	"github.com/alexandreLamarre/ro"
	"github.com/stretchr/testify/assert"
)

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
}
