package dp

import "github.com/hxzqlh/leetcode/util"

// dp[i][j]：以i-1为结尾的字符串word1，和以j-1位结尾的字符串word2，想要达到相等，所需要删除元素的最少次数。
//
// if (word1[i - 1] == word2[j - 1])
// 		两者再做删除操作: dp[i - 1][j - 1]
//
// if (word1[i - 1] != word2[j - 1])
// 		删 word1：dp[i - 1][j]+1
//		删 word2：dp[i][j-1]+1
//		同时删word1[i - 1]和word2[j - 1]：dp[i - 1][j - 1] + 2
//
func MinDelOperations(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	var i, j int
	for i = 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for j = 0; j < n+1; j++ {
		dp[0][j] = j
	}

	for i = 1; i < m+1; i++ {
		for j = 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = util.Min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+2)
			}
		}
	}
	return dp[m][n]
}
