package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortSquares(t *testing.T) {
	cases := []struct {
		in []int
		out []int
	}{
		{[]int{-4,-1,0,3,10}, []int{0,1,9,16,100}},
		{[]int{-7,-3,2,3,11},[]int{4,9,9,49,121}},
	}
	for _, v := range cases {
		assert.Equal(t, SortSquares(v.in), v.out)
	}
}
