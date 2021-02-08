package leetcode

import "testing"

func Test162(t *testing.T) {
	var (
		nums    = []int{1, 2, 1, 3, 5, 6, 4}
		result1 = 1
		result2 = 5
	)
	r := findPeakElement(nums)
	if r != result1 && r != result2 {
		t.Fatal(r, "!=", result1, "and !=", result2)
	}
	t.Log("成功")
}
