package ro_test

import (
	"testing"

	"github.com/alexandreLamarre/ro"
	"github.com/stretchr/testify/assert"
)

// convertsslices of the form []int{3,0,0} to 300
func digitsToInt(arr []int) int {
	ret := 0
	for _, v := range arr {
		ret = ret*10 + v
	}
	return ret
}

func TestAll(t *testing.T) {
	it := ro.Drop(
		ro.Apply(
			ro.Limit(
				ro.Permutations(
					ro.ToSlice(
						ro.Range(0, 6, 1),
					),
					5,
				),
				5,
			),
			func(perm []int) int {
				return digitsToInt(perm)
			},
		),
		func(i int) bool {
			return i > 10000
		},
	)
	res := []int{}
	for v := range it {
		res = append(res, v)
	}
	assert.Equal(
		t,
		[]int{
			1234,
			2134,
		},
		res,
	)
}
