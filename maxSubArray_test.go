package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			[]int{-1, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			[]int{1, 2, 3, 4, -1, 5, 6, 7},
			27,
		},
		{
			[]int{-1},
			-1,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, maxSubArray(testItem.nums))

	}

}
