package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ClimbStairsCase struct {
	n   int
	sum int
}

func TestClimbStairs(t *testing.T) {
	cases := []ClimbStairsCase{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}
	for _, v := range cases {
		assert.Equal(t, ClimbStairs(v.n), v.sum)
	}
}
