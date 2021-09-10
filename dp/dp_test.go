package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditDistance(t *testing.T) {
	cases := []struct {
		word1       string
		word2       string
		minDistance int
	}{
		{"ros", "horse", 3},
		{"ros", "", 3},
		{"", "horse", 5},
		{"ro", "hor", 2},
	}
	for _, v := range cases {
		assert.Equal(t, MinDistance(v.word1, v.word2), v.minDistance)
	}
}

func TestDelOperations(t *testing.T) {
	cases := []struct {
		s   string
		t   string
		res int
	}{
		{"sea", "eat", 2},
		{"b", "", 1},
		{"ab", "ba", 2},
		{"abcd", "bfd", 3},
		{"", "", 0},
	}
	for _, v := range cases {
		assert.Equal(t, MinDelOperations(v.s, v.t), v.res)
	}
}

func TestDistinctSequences(t *testing.T) {
	cases := []struct {
		s   string
		t   string
		res int
	}{
		{"rabbbit", "rabbit", 3},
		{"babgbag", "bag", 5},
		{"", "", 1},
	}
	for _, v := range cases {
		assert.Equal(t, NumDistinct(v.s, v.t), v.res)
	}
}

func TestClimbStairs(t *testing.T) {
	cases := []struct {
		n   int
		sum int
	}{
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

func TestSubSequence(t *testing.T) {
	cases := []struct {
		s   string
		t   string
		res bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}
	for _, v := range cases {
		assert.Equal(t, IsSubSequence(v.s, v.t), v.res)
	}
}
