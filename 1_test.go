package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	var data = []struct {
		nums   []int
		target int
		ans    []int
	}{
		{
			nums:   []int{3, 6, 4, 12, 11},
			target: 9,
			ans:    []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			ans:    []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			ans:    []int{0, 1},
		},
		{
			nums:   []int{3, 3},
			target: 7,
			ans:    nil,
		},
	}

	for _, v := range data {
		assert.ElementsMatch(t, twoSum(v.nums, v.target), v.ans)
	}
}
