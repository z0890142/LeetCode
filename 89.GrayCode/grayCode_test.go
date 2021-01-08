package GrayCode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_grayCode(t *testing.T) {
	tests := []struct {
		n      int
		result []int
	}{
		{
			2,
			[]int{0, 1, 3, 2},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, grayCode(testItem.n))

	}
}
