package PascalsTriangleII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getRow(t *testing.T) {
	tests := []struct {
		rowIndex int
		result   []int
	}{
		{
			3,
			[]int{1, 2, 1},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, getRow(testItem.rowIndex))

	}
}
