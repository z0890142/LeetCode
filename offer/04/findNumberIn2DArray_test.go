package findNumberIn2DArray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findNumberIn2DArray(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		result bool
	}{
		{
			[][]int{[]int{5}, []int{6}},
			6,
			true,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, findNumberIn2DArray(testItem.matrix, testItem.target))

	}
}
