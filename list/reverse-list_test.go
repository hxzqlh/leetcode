package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ReverseCase struct {
	n1 []int
	n2 []int
}

func TestReverse(t *testing.T) {
	cases := []ReverseCase{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{}, []int{}},
	}
	for _, v := range cases {
		l1 := MakeList(v.n1)
		l2 := MakeList(v.n2)
		PrintList(l1)
		r1 := reverseList2(l1)
		PrintList(r1)
		assert.Equal(t, IsEqual(r1, l2), true)
	}
}
