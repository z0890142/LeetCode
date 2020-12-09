package PascalsTriangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate(t *testing.T) {
	tests := []struct {
		numRows int
		result  [][]int
	}{
		{
			3,
			[][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
			},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, generate(testItem.numRows))

	}
}
