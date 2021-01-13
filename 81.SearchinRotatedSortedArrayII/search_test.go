package SearchinRotatedSortedArrayII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result bool
	}{

		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
			true,
		},
		{
			[]int{1, 1, 1, 1, 3},
			3,
			true,
		},
		{
			[]int{1},
			1,
			true,
		},
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			0,
			true,
		},
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			3,
			false,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, search(testItem.nums, testItem.target))

	}

}
