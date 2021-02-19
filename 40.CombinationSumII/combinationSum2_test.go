package CombinationSumII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum2(t *testing.T) {
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
			[]int{10, 1, 2, 7, 6, 1, 5},
			8,
			[][]int{
				[]int{1, 7},
				[]int{1, 2, 5},
				[]int{2, 6},
				[]int{1, 1, 6},
			},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, combinationSum2(testItem.candidates, testItem.target))

	}

}
