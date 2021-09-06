package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MergeCase struct {
	n1 []int
	n2 []int
	n3 []int
}

func TestMerge(t *testing.T) {
	cases := []MergeCase{
		{[]int{1, 2}, []int{2, 3, 5}, []int{1, 2, 2, 3, 5}},
		{[]int{1}, []int{2}, []int{1, 2}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{}, []int{}, []int{}},
	}

	for _, v := range cases {
		l1 := MakeList(v.n1)
		l2 := MakeList(v.n2)
		l3 := MakeList(v.n3)
		ret := mergeList(l1, l2)
		PrintList(l3)
		PrintList(ret)
		assert.Equal(t, IsEqual(ret, l3), true)
	}
}
