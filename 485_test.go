package leetcode

import "testing"

func Test485(t *testing.T) {
	var (
		nums   = []int{1, 1, 0, 1, 1, 1}
		result = 3
	)
	if r := findMaxConsecutiveOnes(nums); r != result {
		t.Fatal(r)
	}
}
