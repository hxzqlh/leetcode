package array

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
// 示例 1： 输入：nums = [-4,-1,0,3,10] 输出：[0,1,9,16,100]
// 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]
//
// 示例 2： 输入：nums = [-7,-3,2,3,11] 输出：[4,9,9,49,121]
//
func SortSquares(arr []int) []int {
	ret := make([]int, len(arr))

	// 绝对值最大的只能是在数组的两端
	pos := len(arr)-1
	i, j := 0, len(arr)-1
	for i<=j && pos >= 0 {
		if arr[i]*arr[i] > arr[j]*arr[j] {
			ret[pos] = arr[i]*arr[i]
			i++
		} else {
			ret[pos] = arr[j]*arr[j]
			j--
		}
		pos--
	}

	return ret
}
