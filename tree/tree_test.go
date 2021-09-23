package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInOrderTraversal(t *testing.T) {
	cases := []struct {
		arr []int
		out []int
	}{
		//      1
		//   2     3
		// 4   5
		{[]int{1, 2, 3, 4, 5}, []int{4, 2, 5, 1, 3}},

		//      1
		//   2      3
		// 4   5  6   7
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7}},

		//      1
		//   2      3
		// 4   5  6
		{[]int{1, 2, 3, 4, 5, 6}, []int{4, 2, 5, 1, 6, 3}},

		//        1
		//     2      3
		// nil   5  6
		{[]int{1, 2, 3, Null, 5, 6}, []int{2, 5, 1, 6, 3}},

		//      1
		//   2        3
		// 4   nil  6
		{[]int{1, 2, 3, 4, Null, 6}, []int{4, 2, 1, 6, 3}},

		//        1
		//     2     3
		// 	 4   5  6  7
		// 8
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{8, 4, 2, 5, 1, 6, 3, 7}},

		//        1
		//     2     3
		// 	 4   5  6  nil
		// 8
		{[]int{1, 2, 3, 4, 5, 6, Null, 8}, []int{8, 4, 2, 5, 1, 6, 3}},

		//      1
		//   2    nil
		// 4
		{[]int{1, 2, Null, 4}, []int{4, 2, 1}},

		//         1
		//    nil       3
		// nil   nil  6   7
		{[]int{1, Null, 3, Null, Null, 6, 7}, []int{1, 6, 3, 7}},
	}

	for _, v := range cases {
		tree := Arr2Tree(v.arr)
		assert.Equal(t, InOrderTraversal3(tree), v.out)
	}
}
