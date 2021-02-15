package leetcode

import "testing"

func Test162(t *testing.T) {
	var (
		nums1    = []int{1, 2, 1, 3, 5, 6, 4}
		result11 = 1
		result12 = 5
	)
	if r := findPeakElement(nums1); r != result11 && r != result12 {
		t.Fatal(r, "!=", result11, "and !=", result12)
	}

	var (
		nums2   = []int{1, 2}
		result2 = 1
	)
	if r := findPeakElement(nums2); r != result2 {
		t.Fatal(r, "!=", result2)
	}
}
