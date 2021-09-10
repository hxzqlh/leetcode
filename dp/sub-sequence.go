package dp

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。
// （例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
//
// 示例 1： 输入：s = "abc", t = "ahbgdc" 输出：true
// 示例 2： 输入：s = "axc", t = "ahbgdc" 输出：false
// 示例 3： 输入：s = "", t = "ahbgdc" 输出：true

// 双指针法
func IsSubSequence2(s string, t string) bool {
	p, size := 0, len(s)
	for i := 0; i < len(t); i++ {
		if p < size && s[p] == t[i] {
			p++
		}
	}
	return p == size
}

// dp[i][j] 表示以下标i-1为结尾的字符串s，和以下标j-1为结尾的字符串t，相同子序列的长度为dp[i][j]。
//
// if (s[i - 1] == t[j - 1])
// 		找到了一个相同的字符，相同子序列长度自然要在dp[i-1][j-1]的基础上加1, 即：dp[i][j] = dp[i - 1][j - 1] + 1
// if (s[i - 1] != t[j - 1])
//		t要删除元素，继续匹配, 即：dp[i][j] = dp[i][j - 1]
func IsSubSequence(s string, t string) bool {
	m, n := len(s), len(t)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	var i, j int
	for i = 0; i < m+1; i++ {
		dp[i][0] = 0
	}
	for j = 0; j < n+1; j++ {
		dp[0][j] = 0
	}

	for i = 1; i < m+1; i++ {
		for j = 1; j < n+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n] == m
}
