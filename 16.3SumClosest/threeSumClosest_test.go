package threeSumClosest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSumClosest(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result int
	}{
		{
			[]int{-1, 2, 1, -4},
			1,
			2,
		},
		{
			[]int{1, 1, 1, 0},
			100,
			3,
		},
		{
			[]int{-1, 1, 1, 0},
			-4,
			0,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, threeSumClosest(testItem.nums, testItem.target))

	}

}
