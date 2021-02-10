package SpiralMatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_spiralOrder(t *testing.T) {
	tests := []struct {
		matrix [][]int
		result []int
	}{
		{
			[][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}, []int{13, 14, 15, 16}},
			[]int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
		{
			[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			[][]int{[]int{1}, []int{5}},
			[]int{1, 5},
		},
		{
			[][]int{[]int{7}, []int{9}, []int{6}},
			[]int{7, 9, 6},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, spiralOrder(testItem.matrix))

	}
}
