package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxChunksToSorted(t *testing.T) {
	tests := []struct {
		arr    []int
		result int
	}{
		{
			[]int{4, 3, 2, 1, 0},
			1,
		},
		{
			[]int{1, 0, 2, 3, 4},
			4,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, maxChunksToSorted(testItem.arr))
	}

}
