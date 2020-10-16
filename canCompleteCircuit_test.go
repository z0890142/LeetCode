package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canCompleteCircuit(t *testing.T) {
	tests := []struct {
		gas    []int
		cost   []int
		result int
	}{
		{
			[]int{3, 3, 4},
			[]int{3, 4, 4},
			-1,
		},
		{
			[]int{5, 1, 2, 3, 4},
			[]int{4, 4, 1, 5, 1},
			4,
		},

		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			3,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, canCompleteCircuit(testItem.gas, testItem.cost))

	}

}
