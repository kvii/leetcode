package leetcode

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) (ans string) {
	dp := longestPalindromeDP(s)

	for i := range s {
		for j := i; j < len(s); j++ {
			if dp[i][j] && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}

	return
}

func longestPalindromeDP(s string) [][]bool {
	// make a n*n matrix
	dp := make([][]bool, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true

		// 只计算矩阵的上半部分
		for j := i + 1; j < len(s); j++ {
			// i+1 >= j-1 规避矩阵下半部分没算导致的问题
			dp[i][j] = s[i] == s[j] && (i+1 >= j-1 || dp[i+1][j-1])
		}
	}

	return dp
}

// @lc code=end
