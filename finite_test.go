package ro_test

import (
	"testing"

	"github.com/alexandreLamarre/ro"
	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	r1 := []int{}
	for v := range ro.Range(1, 10, 2) {
		r1 = append(r1, v)
	}
	assert.Equal(t, []int{1, 3, 5, 7, 9}, r1)

	r2 := []int{}
	for v := range ro.Range(1, 6, 0) {
		r2 = append(r2, v)
	}
	assert.Equal(t, []int{}, r2)

	r3 := []int{}
	for v := range ro.Range(1, 1, 1) {
		r3 = append(r3, v)
	}
	assert.Equal(t, []int{}, r3)

	r4 := []int8{}
	for v := range ro.Range(int8(1), int8(10), int8(2)) {
		r4 = append(r4, v)
	}
	assert.Equal(t, []int8{1, 3, 5, 7, 9}, r4)
}
