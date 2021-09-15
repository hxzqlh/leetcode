package interval

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	cases := []struct{
		intervals [][]int
		newInterval []int
		res [][]int
	}{
		{[][]int{{1,3},{6,9}}, []int{2,5}, [][]int{{1,5},{6,9}}},
		{[][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}, []int{4,8}, [][]int{{1,2},{3,10},{12,16}}},
		{[][]int{},[]int{5,7},[][]int{{5,7}} },
		{[][]int{{1,5}}, []int{2,3}, [][]int{{1,5}}},
		{[][]int{{1,5}}, []int{2,7}, [][]int{{1,7}}},
		{[][]int{{5,7}}, []int{2,3}, [][]int{{2,3},{5,7}}},
	}
	for _, v := range cases {
		assert.Equal(t, Insert(v.intervals, v.newInterval), v.res)
	}
}
