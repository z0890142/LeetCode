package Combinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combine(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		result [][]int
	}{
		// {
		// 	4,
		// 	2,
		// 	[][]int{[]int{1, 2}, []int{1, 3}, []int{1, 4}, []int{2, 3}, []int{2, 4}, []int{3, 4}},
		// },
		{
			5,
			4,
			[][]int{[]int{1, 2, 3, 4}, []int{1, 2, 3, 5}, []int{1, 2, 4, 5}, []int{1, 3, 4, 5}, []int{2, 3, 4, 5}},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, combine(testItem.n, testItem.k))

	}

}
