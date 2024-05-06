//go:build goexperiment.rangefunc

package parallel_test

import (
	"testing"
	"time"

	pro "github.com/alexandreLamarre/ro/parallel"
	"github.com/stretchr/testify/assert"
)

func TestPool(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res := make([]int, 0, 3)

	for x := range pro.Pool(xs, 3) {
		res = append(res, x)
		time.Sleep(time.Millisecond)
		break
	}

	assert.Len(t, res, 3)
}
