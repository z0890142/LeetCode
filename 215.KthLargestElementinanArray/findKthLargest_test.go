package KthLargestElementinanArray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findKthLargest(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		result int
	}{
		{
			[]int{3, 2, 1, 5, 6, 4},
			2,
			5,
		},
		// {
		// 	[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		// 	4,
		// 	4,
		// },
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, findKthLargest(testItem.nums, testItem.k))
	}

}
