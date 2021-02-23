package NextPermutation

import (
	"testing"
)

func Test_search(t *testing.T) {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			[]int{1, 2, 3},
			0,
		},
		{
			[]int{1, 5, 1},
			3,
		},
	}

	for _, testItem := range tests {
		nextPermutation(testItem.nums)

	}

}
