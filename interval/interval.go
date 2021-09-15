package interval

import "github.com/hxzqlh/leetcode/util"

// 给你一个无重叠的，按照区间起始端点排序的区间列表。
// 在列表中插入一个新的非空区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
//
// 输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
// 输出：[[1,5],[6,9]]
//
// 输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// 输出：[[1,2],[3,10],[12,16]]
// 解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
//
// 输入：intervals = [], newInterval = [5,7]
// 输出：[[5,7]]
//
// 输入：intervals = [[1,5]], newInterval = [2,3]
// 输出：[[1,5]]
//
// 输入：intervals = [[1,5]], newInterval = [2,7]
// 输出：[[1,7]]
//
// 输入：intervals = [[5,7]], newInterval = [2,3]
// 输出：[[2,3],[5,7]]
func Insert(intervals [][]int, newInterval []int) [][]int {
	ret := [][]int{}
	left, right := newInterval[0], newInterval[1]

	var i int
	// left > intervals[i][1]
	for ; i < len(intervals) &&  left > intervals[i][1]; i++{
		ret = append(ret, intervals[i])
	}

	// already known: left <= intervals[i][1]
	// find: right >= intervals[i][0]
	for ; i < len(intervals) && right >= intervals[i][0]; i++{
		left = util.Min(intervals[i][0], left)
		right = util.Max(intervals[i][1], right)
	}
	ret = append(ret, []int{left, right})

	// right < intervals[i][0]
	for ; i < len(intervals) && right < intervals[i][0]; i++{
		ret = append(ret, intervals[i])
	}

	return ret
}
