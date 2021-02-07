package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for j, num := range nums {
		if i, ok := mp[num]; ok {
			return []int{i, j}
		}
		mp[target-num] = j
	}
	return nil
}

// @lc code=end
