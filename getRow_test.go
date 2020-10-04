package main

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
			2,
			[]int{1, 2, 2},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, getRow(testItem.rowIndex))

	}
}
