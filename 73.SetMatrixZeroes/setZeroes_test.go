package SetMatrixZeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setZeroes(t *testing.T) {
	tests := []struct {
		matrix [][]int
		result [][]int
	}{
		{
			[][]int{
				[]int{1, 1, 1},
				[]int{1, 0, 1},
				[]int{1, 1, 1},
			},
			[][]int{
				[]int{1, 0, 1},
				[]int{0, 0, 0},
				[]int{1, 0, 1},
			},
		},
		{
			[][]int{
				[]int{0, 1, 2, 0},
				[]int{3, 4, 5, 2},
				[]int{1, 3, 1, 5},
			},
			[][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 4, 5, 0},
				[]int{0, 3, 1, 0},
			},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, setZeroes(testItem.matrix))

	}
}
