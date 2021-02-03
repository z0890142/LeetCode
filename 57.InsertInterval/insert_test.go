package InsertInterval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insert(t *testing.T) {
	tests := []struct {
		intervals   [][]int
		newInterval []int
		result      [][]int
	}{
		{
			[][]int{[]int{1, 3}, []int{6, 9}},
			[]int{2, 5},
			[][]int{[]int{1, 5}, []int{6, 9}},
		},

		{
			[][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}},
			[]int{4, 8},
			[][]int{[]int{1, 2}, []int{3, 10}, []int{12, 16}},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, insert(testItem.intervals, testItem.newInterval))
	}

}
