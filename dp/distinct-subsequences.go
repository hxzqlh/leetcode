package dp

// 给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
//
// 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。
// 例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是
//
// dp[i][j]：以i-1为结尾的s子序列中出现以j-1为结尾的t的个数为dp[i][j]
//
// if (s[i - 1] == t[j - 1])
// 		1. 用 s[i - 1] 来匹配: dp[i - 1][j - 1]
// 		2. 不用 s[i - 1] 来匹配: dp[i - 1][j]
//
// 		如 s：bagg 和 t：bag ，s[3] 和 t[2]是相同的
// 		1. s用s[3]来匹配，即：s[0]s[1]s[3]组成的bag
// 		2. s不用s[3]来匹配，即用s[0]s[1]s[2]组成的bag
//
// if s[i - 1] != t[j - 1]
// 		不用s[i - 1]来匹配：dp[i - 1][j]
func NumDistinct(s string, t string) int {
	m, n := len(s), len(t)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	var i, j int
	for i = 1; i < m+1; i++ {
		dp[i][0] = 1
	}
	for j = 1; j < n+1; j++ {
		dp[0][j] = 0
	}
	dp[0][0] = 1

	for i = 1; i < m+1; i++ {
		for j = 1; j < n+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}
