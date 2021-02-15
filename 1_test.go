package leetcode

import (
	"testing"
)

func Test1(t *testing.T) {
	var (
		nums   = []int{3, 6, 4, 12, 11}
		target = 9
		result = []int{0, 1}
	)
	arr := twoSum(nums, target)

	if err := intSlice(arr, result); err != nil {
		t.Fatal(err)
	}
}
