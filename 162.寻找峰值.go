package leetcode

/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
func findPeakElement(nums []int) int {
	// A 的 2..(n-1) 中 A[i-1] < A[i] > A[i+1] 的位置
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			return i
		}
	}
	return 0
}

// @lc code=end
