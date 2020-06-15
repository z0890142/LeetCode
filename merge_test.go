package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		result    [][]int
	}{
		// {
		// 	[][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}},
		// 	[][]int{[]int{1, 6}, []int{8, 10}, []int{15, 18}},
		// },
		{
			[][]int{[]int{4, 5}, []int{1, 4}},
			[][]int{[]int{1, 5}},
		},
		// {
		// 	[][]int{[]int{1, 4}, []int{5, 6}, []int{7, 8}},
		// 	[][]int{[]int{1, 4}, []int{5, 6}, []int{7, 8}},
		// },
		// {
		// 	[][]int{[]int{1, 4}, []int{2, 3}},
		// 	[][]int{[]int{1, 4}},
		// },
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, merge(testItem.intervals))

	}

}
