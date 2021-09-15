package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	cases := []struct{
		in []int
		out [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2},{-1, 0, 1}}},
	}
	for _, v := range cases {
		assert.Equal(t, v.out, ThreeSum(v.in))
	}
}
