package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateMatrix(t *testing.T) {
	tests := []struct {
		n      int
		result [][]int
	}{
		{
			3,
			[][]int{
				[]int{1, 2, 3},
				[]int{8, 9, 4},
				[]int{7, 6, 5},
			},
		},
		{
			5,
			[][]int{
				[]int{1, 2, 3, 4, 5},
				[]int{16, 17, 18, 19, 6},
				[]int{15, 24, 25, 20, 7},
				[]int{14, 23, 22, 21, 8},
				[]int{13, 12, 11, 10, 9},
			},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, generateMatrix(testItem.n))

	}

}
