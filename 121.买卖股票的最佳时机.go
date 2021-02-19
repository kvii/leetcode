package leetcode

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {

	var (
		min  = prices[0] // A[1..i-1] 的最小值
		prev int         // A[1..i-1] 的最大收益
	)
	for i := 1; i < len(prices); i++ {

		// 跌停板 -> 更新下标
		if prices[i] < min {
			min = prices[i]
			continue
		}
		// 收益更高 -> 更新收益
		if curr := prices[i] - min; curr > prev {
			prev = curr
		}
	}

	return prev
}

// @lc code=end
