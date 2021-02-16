package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test28(t *testing.T) {
	var data = []struct {
		a, b string
		ans  int
	}{
		{"", "", 0},
		{"hello", "ll", 2},
		{"golang", "ll", -1},
		{"abc", "c", 2},
	}

	for _, v := range data {
		assert.Equal(t, strStr(v.a, v.b), v.ans)
	}
}
