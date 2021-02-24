package RemoveElement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		nums   []int
		val    int
		result int
	}{
		{
			[]int{3, 2, 2, 3},
			3,
			2,
		},
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			5,
		},
		{
			[]int{},
			0,
			0,
		},
		{
			[]int{0},
			1,
			1,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, removeElement(testItem.nums, testItem.val))

	}
}
