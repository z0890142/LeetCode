package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canJump(t *testing.T) {
	tests := []struct {
		nums   []int
		result bool
	}{
		{
			[]int{2, 3, 1, 1, 4},
			true,
		},
		{
			[]int{3, 2, 1, 0, 4},
			false,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, canJump(testItem.nums))

	}

}
