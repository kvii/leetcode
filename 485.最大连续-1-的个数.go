package leetcode

/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	return findMaxSubarray(nums, 0, len(nums)-1)
}

func findMaxSubarray(A []int, low, high int) int {

	if low == high {
		return A[low] // A[low] == 1 ? 1 : 0
	}

	mid := (low + high) / 2
	var (
		max1 = findMaxSubarray(A, low, mid)
		max2 = findMaxSubarray(A, mid+1, high)
		max3 = findMaxCrossingSubarray(A, low, mid, high)
	)

	if max1 >= max2 && max1 >= max3 {
		return max1
	}
	if max2 >= max1 && max2 >= max3 {
		return max2
	}
	return max3
}

func findMaxCrossingSubarray(A []int, low, mid, high int) (count int) {
	for i := mid; i >= low; i-- {
		if A[i] == 0 {
			break
		}
		count++
	}

	for i := mid + 1; i <= high; i++ {
		if A[i] == 0 {
			break
		}
		count++
	}

	return
}

func findMaxSubarray2(A []int) (ans int) {

	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			continue
		}

		var j, count int
		for j = i; j < len(A); j++ {
			if A[j] == 0 {
				break
			}
			count++
		}

		if count > ans {
			ans = count
		}
		i = j
	}
	return
}

// @lc code=end
