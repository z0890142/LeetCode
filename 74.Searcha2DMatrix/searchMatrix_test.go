package Searcha2DMatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		result bool
	}{
		{
			[][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}},
			3,
			true,
		},
		{
			[][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}},
			13,
			false,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, searchMatrix(testItem.matrix, testItem.target))

	}

}
