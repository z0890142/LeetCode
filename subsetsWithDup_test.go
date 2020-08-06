package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsetsWithDupHelper(t *testing.T) {
	tests := []struct {
		nums   []int
		result [][]int
	}{
		{
			[]int{1, 2, 2},
			[][]int{[]int{}, []int{1}, []int{1, 2}, []int{1, 2, 3}, []int{2}, []int{2, 2}},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, subsetsWithDup(testItem.nums))

	}

}
