package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	data := []struct {
		x, want int
	}{
		{x: 123, want: 321},
		{x: 120, want: 21},
		{x: 0, want: 0},
		{x: -123, want: -321},
		{x: 2<<31 - 1, want: 0},
		{x: -2 << 31, want: 0},
	}

	for i, v := range data {
		assert.Equal(t, v.want, reverse(v.x), "%d x:%d y:%d", i, v.want)
	}
}
