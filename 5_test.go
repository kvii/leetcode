package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test5(t *testing.T) {
	data := []struct {
		s   string
		ans []string
	}{
		{"babad", []string{"bab", "aba"}},
		{"cbbd", []string{"bb"}},
		{"a", []string{"a"}},
		{"ac", []string{"a", "c"}},
		{"abcba", []string{"abcba"}},
	}

	for _, v := range data {
		assert.Contains(t, v.ans, longestPalindrome(v.s))
	}
}
