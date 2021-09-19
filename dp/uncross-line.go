package dp

import "github.com/hxzqlh/leetcode/util"

// MaxUncrossedLines
// 输入：nums1 = [1,4,2], nums2 = [1,2,4]
// 输出：2
// 解释：可以画出两条不交叉的线: 1-1, 4-4。
// 但无法画出第三条不相交的直线，因为从 nums1[1]=4 到 nums2[2]=4 的直线将与从 nums1[2]=2 到 nums2[1]=2 的直线相交。
//
// 分析：绘制一些连接两个数字 A[i] 和 B[j] 的直线，只要 A[i] == B[j]，且直线不能相交！
// 直线不能相交，这就是说明在字符串A中找到一个与字符串B相同的子序列，且这个子序列不能改变相对顺序，
// 只要相对顺序不改变，链接相同数字的直线就不会相交。
//
// 其实也就是说A和B的最长公共子序列是[1,4]，长度为2。
// 这个公共子序列指的是相对顺序不变（即数字4在字符串A中数字1的后面，那么数字4也应该在字符串B数字1的后面）
// 本题说是求绘制的最大连线数，其实就是: 求两个字符串的最长公共子序列的长度！
func MaxUncrossedLines(nums1 []int, nums2 []int) int {
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
			} else {
				dp[i][j] = util.Max(dp[i-1][j], dp[i][j-1])
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res
}
