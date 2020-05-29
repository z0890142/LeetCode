package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fourSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result [][]int
	}{
		{
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}},
		},
		{
			[]int{0, 0, 0, 0},
			1,
			[][]int{},
		},
		{
			[]int{-3, -1, 0, 2, 4, 5},
			2,
			[][]int{[]int{-3, -1, 2, 4}},
		},

		{
			[]int{-3, -2, -1, 0, 0, 1, 2, 3},
			0,
			[][]int{[]int{-3, -2, 2, 3}, []int{-3, -1, 1, 3}, []int{-3, 0, 0, 3}, []int{-3, 0, 1, 2}, []int{-2, -1, 0, 3}, []int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, fourSum(testItem.nums, testItem.target))

	}

}
