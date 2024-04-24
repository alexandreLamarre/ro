//go:build goexperiment.rangefunc

package ro_test

import (
	"testing"

	"github.com/alexandreLamarre/ro"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.Zip([]int{1, 2, 3}, []int{4, 5, 6}) {
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
}

func TestZipIter(t *testing.T) {
	res := []lo.Tuple2[int, int]{}
	for v := range ro.ZipIter(ro.SeqAsIter([]int{1, 2, 3}), ro.SeqAsIter([]int{4, 5, 6})) {
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
}
