package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type AddCase struct {
	n1 []int
	n2 []int
	n3 []int
}

func TestAddTwoNums2(t *testing.T) {
	cases := []AddCase{
		{[]int{6, 1, 7}, []int{2, 9, 5}, []int{9, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{2, 4, 6}, []int{1, 2, 5, 9, 1}},
		{[]int{9, 1, 5}, []int{1, 6, 5}, []int{1, 0, 8, 0}},
		{[]int{9, 9, 9}, []int{1}, []int{1, 0, 0, 0}},
	}

	for _, v := range cases {
		l1 := MakeList(v.n1)
		l2 := MakeList(v.n2)
		l3 := MakeList(v.n3)
		ret := addTwoNumbers2(l1, l2)
		PrintList(l3)
		PrintList(ret)
		assert.Equal(t, IsEqual(ret, l3), true)
	}
}
