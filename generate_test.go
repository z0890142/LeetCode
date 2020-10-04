package main

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
			5,
			[][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
				[]int{1, 3, 3, 1},
				[]int{1, 4, 6, 4, 1},
			},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, generate(testItem.numRows))

	}

}
