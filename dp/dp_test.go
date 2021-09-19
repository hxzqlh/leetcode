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

func TestCommonLength(t *testing.T) {
	cases := []struct {
		nums1 []int
		nums2 []int
		ret int
	}{
		{[]int{1,2,3,2,1}, []int{3,2,1,4,7}, 3},
	}
	for _, v := range cases {
		assert.Equal(t, FindLength(v.nums1, v.nums2), v.ret)
	}
}

func TestCommonSubsequence(t *testing.T) {
	cases := []struct {
		s   string
		t   string
		res int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"abc", "", 0},
		{"", "def", 0},
	}
	for _, v := range cases {
		assert.Equal(t, LongestCommonSubsequence(v.s, v.t), v.res)
	}
}

func TestUncrossedLines(t *testing.T) {
	cases := []struct {
		nums1   []int
		nums2   []int
		res int
	}{
		{[]int{1,4,2}, []int{1,2,4},2},
		{[]int{1,4,2}, []int{1,4,2},3},
		{[]int{2,5,1,2,5}, []int{10,5,2,1,5,2},3},
		{[]int{1,3,7,1,7,5}, []int{1,9,2,5,1},2},
	}
	for _, v := range cases {
		assert.Equal(t, MaxUncrossedLines(v.nums1, v.nums2), v.res)
	}
}