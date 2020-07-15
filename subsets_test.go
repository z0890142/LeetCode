package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsets(t *testing.T) {
	tests := []struct {
		nums   []int
		result [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{[]int{}, []int{1}, []int{2}, []int{3}, []int{1, 2, 3}, []int{1, 2}, []int{1, 3}, []int{2, 3}},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, subsets(testItem.nums))

	}

}
