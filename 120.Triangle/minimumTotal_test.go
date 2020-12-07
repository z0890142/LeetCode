package Triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumTotal(t *testing.T) {
	tests := []struct {
		triangle [][]int
		result   int
	}{
		{
			[][]int{
				[]int{2},
				[]int{3, 4},
				[]int{6, 5, 7},
				[]int{4, 1, 8, 3},
			},
			11,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, minimumTotal(testItem.triangle))
	}

}
