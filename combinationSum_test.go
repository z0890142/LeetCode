package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		result     [][]int
	}{
		// {
		// 	[]int{2, 3, 6, 7},
		// 	7,
		// 	[][]int{
		// 		[]int{7},
		// 		[]int{2, 2, 3},
		// 	},
		// },
		{
			[]int{2, 3, 5},
			8,
			[][]int{
				[]int{2, 2, 2, 2},
				[]int{2, 3, 3},
				[]int{3, 5},
			},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, combinationSum(testItem.candidates, testItem.target))

	}

}
