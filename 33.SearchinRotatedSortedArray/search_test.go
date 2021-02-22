package SearchinRotatedSortedArray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result int
	}{
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
			4,
		},
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			3,
			-1,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, search(testItem.nums, testItem.target))

	}

}
