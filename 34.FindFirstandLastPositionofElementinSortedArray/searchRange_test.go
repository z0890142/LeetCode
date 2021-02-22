package FindFirstandLastPositionofElementinSortedArray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result []int
	}{
		{
			[]int{5, 7, 7, 8, 8, 10},
			8,
			[]int{3, 4},
		},
		{
			[]int{5, 7, 7, 8, 8, 10},
			6,
			[]int{-1, -1},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, searchRange(testItem.nums, testItem.target))

	}

}
