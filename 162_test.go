package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test162(t *testing.T) {
	var data = []struct {
		nums, ans []int
	}{
		{
			nums: []int{1, 2, 1, 3, 5, 6, 4},
			ans:  []int{1, 5},
		},
		{
			nums: []int{1, 2},
			ans:  []int{1},
		},
		{
			nums: []int{1},
			ans:  []int{0},
		},
		{
			nums: []int{2, 1},
			ans:  []int{0},
		},
	}

	for _, v := range data {
		assert.Contains(t, v.ans, findPeakElement(v.nums))
	}
}
