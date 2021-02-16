package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test485(t *testing.T) {
	var data = []struct {
		nums []int
		ans  int
	}{
		{ans: 3, nums: []int{1, 1, 0, 1, 1, 1}},
		{ans: 0, nums: []int{0}},
		{ans: 1, nums: []int{1}},
	}

	for _, v := range data {
		assert.Equal(t, findMaxConsecutiveOnes(v.nums), v.ans)
	}
}
