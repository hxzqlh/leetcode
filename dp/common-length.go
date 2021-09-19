package dp

// FindLength
// 给两个整数数组，返回两个数组中公共的、长度最长的子数组的长度。
//
// 输入：
// A: [1,2,3,2,1]
// B: [3,2,1,4,7]
// 输出：3
// 解释：
// 长度最长的公共子数组是 [3, 2, 1] 。
func FindLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	res := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++{
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1]+1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res
}
