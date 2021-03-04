package fourSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result [][]int
	}{
		{
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{
				[]int{-2, -1, 1, 2},
				[]int{-2, 0, 0, 2},
				[]int{-1, 0, 0, 1},
			},
		},
	}
	for _, testItem := range tests {
		assert.Equal(t, testItem.result, fourSum(testItem.nums, testItem.target))

	}

}
