package MaximumProductSubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		nums   []int
		result int
	}{
		// {
		// 	"a good   example",
		// 	"example good a",
		// },

		{
			[]int{-2, -3, 7},
			42,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, maxProduct(testItem.nums))
	}

}
