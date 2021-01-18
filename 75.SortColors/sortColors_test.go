package SortColors

import (
	"testing"
)

func Test_sortColors(t *testing.T) {
	tests := []struct {
		nums   []int
		result []int
	}{

		{
			[]int{2, 0, 2, 1, 1, 0},
			[]int{0, 0, 1, 1, 2, 2},
		},
	}

	for _, testItem := range tests {
		sortColors(testItem.nums)

	}

}
