package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums []int

		result int
		check  []int
	}{
		{
			[]int{0, 1, 1, 1, 2, 2, 3, 3, 4, 5},
			6,
			[]int{0, 1, 2, 3, 4, 5},
		},
		{
			[]int{0},
			1,
			[]int{0},
		},
		{
			[]int{},
			0,
			[]int{},
		},
		{
			[]int{0, 1, 2, 3, 4},
			5,
			[]int{0, 1, 2, 3, 4},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, removeDuplicates(testItem.nums))

	}
}
