package FindPeakElement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPeakElement(t *testing.T) {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			[]int{1, 3, 2, 1},
			1,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, findPeakElement(testItem.nums))
	}

}
