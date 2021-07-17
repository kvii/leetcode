package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {

	var matrix = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	for _, row := range matrix {
		for _, target := range row {
			assert.True(t, searchMatrix(matrix, target))
		}
	}

	assert.False(t, searchMatrix(matrix, 0))
	assert.False(t, searchMatrix(matrix, 200))
}
