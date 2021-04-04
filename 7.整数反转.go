package leetcode

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) (y int) {
	for ; x != 0; x /= 10 {
		y = y*10 + x%10
	}
	
	if int(int32(y)) != y {
		return 0
	}

	return
}

// @lc code=end
