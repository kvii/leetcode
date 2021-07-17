package leetcode

func maximumPopulation(logs [][]int) int {
	const (
		offset = 1950
		years  = 2050 - 1950 + 1
	)

	// 人口变化量
	var delta [years]int
	for i := range logs {
		birth := logs[i][0] - offset
		death := logs[i][1] - offset

		delta[birth]++
		delta[death]--
	}

	// 最大人口 下标 当前人口
	var max, ans, cur int
	for i := 0; i < years; i++ {
		if cur += delta[i]; cur > max {
			max = cur
			ans = i
		}
	}

	return offset + ans
}
