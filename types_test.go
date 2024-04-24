package ro_test

import (
	"testing"

	"github.com/alexandreLamarre/ro"
	"github.com/stretchr/testify/assert"
)

func TestSeqAsIter(t *testing.T) {
	i := ro.SeqAsIter([]string{"a", "b", "c"})
	res := []string{}
	for v := range i {
		res = append(res, v)
	}
	assert.Equal(t, []string{"a", "b", "c"}, res)
}
