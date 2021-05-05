package leetcode

/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	if n < 2 {
		return n
	}

	var a, b = 0, 1
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}

	return a + b
}

// @lc code=end
