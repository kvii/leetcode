package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test485(t *testing.T) {
	data := []struct {
		nums   []int
		result int
	}{
		{result: 3, nums: []int{1, 1, 0, 1, 1, 1}},
		{result: 0, nums: []int{0}},
		{result: 1, nums: []int{1}},
	}

	for _, v := range data {
		assert.Equal(t, findMaxConsecutiveOnes(v.nums), v.result)
	}
}
