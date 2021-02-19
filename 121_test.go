package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test121(t *testing.T) {
	data := []struct {
		prices []int
		ans    int
	}{
		{ans: 5, prices: []int{7, 1, 5, 3, 6, 4}},
		{ans: 0, prices: []int{7, 6, 4, 3, 1}},
		{ans: 0, prices: []int{4}},
		{ans: 0, prices: []int{4, 1}},
		{ans: 3, prices: []int{1, 4}},
	}

	for _, v := range data {
		assert.Equal(t, v.ans, maxProfit(v.prices))
	}
}

func Benchmark121(b *testing.B) {
	// 下标法	249733096	4.80 ns/op
	// range法	195351862	6.28 ns/op

	var prices = []int{7, 1, 5, 3, 6, 4}
	for i := 0; i < b.N; i++ {
		maxProfit(prices)
	}
}
