package RemoveDuplicatesfromSortedArrayII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums []int

		result int
		check  []int
	}{
		{
			[]int{1, 1, 1, 2, 2, 3},
			5,
			[]int{1, 1, 2, 2, 3},
		},
		{
			[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			7,
			[]int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			7,
			[]int{0, 0, 1, 1, 2, 3, 3},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, removeDuplicates(testItem.nums))

	}
}
