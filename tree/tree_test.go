package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInOrderTraversal(t *testing.T) {
	cases := []struct {
		arr  []int
		in   []int
		pre  []int
		post []int
	}{
		//      1
		//   2     3
		// 4   5
		{[]int{1, 2, 3, 4, 5}, []int{4, 2, 5, 1, 3}, []int{1, 2, 4, 5, 3}, []int{4, 5, 2, 3, 1}},

		//      1
		//   2      3
		// 4   5  6   7
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7}, []int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}},

		//      1
		//   2      3
		// 4   5  6
		{[]int{1, 2, 3, 4, 5, 6}, []int{4, 2, 5, 1, 6, 3}, []int{1, 2, 4, 5, 3, 6}, []int{4, 5, 2, 6, 3, 1}},

		//        1
		//     2      3
		// nil   5  6
		{[]int{1, 2, 3, Null, 5, 6}, []int{2, 5, 1, 6, 3}, []int{1, 2, 5, 3, 6}, []int{5, 2, 6, 3, 1}},

		//      1
		//   2        3
		// 4   nil  6
		{[]int{1, 2, 3, 4, Null, 6}, []int{4, 2, 1, 6, 3}, []int{1, 2, 4, 3, 6}, []int{4, 2, 6, 3, 1}},

		//        1
		//     2     3
		// 	 4   5  6  7
		// 8
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{8, 4, 2, 5, 1, 6, 3, 7}, []int{1, 2, 4, 8, 5, 3, 6, 7}, []int{8, 4, 5, 2, 6, 7, 3, 1}},

		//        1
		//     2     3
		// 	 4   5  6  nil
		// 8
		{[]int{1, 2, 3, 4, 5, 6, Null, 8}, []int{8, 4, 2, 5, 1, 6, 3}, []int{1, 2, 4, 8, 5, 3, 6}, []int{8, 4, 5, 2, 6, 3, 1}},

		//      1
		//   2    nil
		// 4
		{[]int{1, 2, Null, 4}, []int{4, 2, 1}, []int{1, 2, 4}, []int{4, 2, 1}},

		//         1
		//    nil       3
		// nil   nil  6   7
		{[]int{1, Null, 3, Null, Null, 6, 7}, []int{1, 6, 3, 7}, []int{1, 3, 6, 7}, []int{6, 7, 3, 1}},
	}

	for _, v := range cases {
		tree := Arr2Tree(v.arr)
		// assert.Equal(t, InOrderTraversal3(tree), v.in)
		// assert.Equal(t, PreOrderTraversal(tree), v.pre)
		assert.Equal(t, PostOrderTraversal(tree), v.post)
	}
}

func TestLevelTraversal(t *testing.T) {
	cases := []struct {
		arr []int
		res [][]int
	}{
		{[]int{3, 9, 20, Null, Null, 15, 7}, [][]int{{3}, {9, 20}, {15, 7}}},
	}
	for _, v := range cases {
		tree := Arr2Tree(v.arr)
		assert.Equal(t, v.res, LevelOrder(tree))
	}
}

func TestRightSideView(t *testing.T) {
	cases := []struct {
		arr []int
		res []int
	}{
		{[]int{1, 2, 3, Null, 5, Null, 4}, []int{1, 3, 4}},
		{[]int{3, 9, 20, Null, Null, 15, 7}, []int{3, 20, 7}},
		{[]int{1}, []int{1}},
	}
	for _, v := range cases {
		tree := Arr2Tree(v.arr)
		assert.Equal(t, v.res, RightSideView(tree))
	}
}
