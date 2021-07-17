package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test5750(t *testing.T) {

	var data = []struct {
		log  [][]int
		want int
	}{
		{
			log: [][]int{
				{1993, 1999},
				{2000, 2010},
			},
			want: 1993,
		},
		{
			log: [][]int{
				{1950, 1961},
				{1960, 1971},
				{1970, 1981},
			},
			want: 1960,
		},
	}

	for _, v := range data {
		assert.Equal(t, v.want, maximumPopulation(v.log))
	}
}
