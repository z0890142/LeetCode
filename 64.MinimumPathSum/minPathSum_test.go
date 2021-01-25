package MinimumPathSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minPathSum(t *testing.T) {

	tests := []struct {
		gird   [][]int
		result int
		Name   string
	}{
		{
			[][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}},
			7,
			"test 1",
		},
	}

	for _, testItem := range tests {

		assert.Equal(t, testItem.result, minPathSum(testItem.gird))

	}
}
