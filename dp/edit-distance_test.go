package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type EditDistanceCase struct {
	word1       string
	word2       string
	minDistance int
}

func TestEditDistance(t *testing.T) {
	cases := []EditDistanceCase{
		{"ros", "horse", 3},
		{"ros", "", 3},
		{"", "horse", 5},
		{"ro", "hor", 2},
	}
	for _, v := range cases {
		assert.Equal(t, MinDistance(v.word1, v.word2), v.minDistance)
	}
}
