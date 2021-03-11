package JumpGameII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jump(t *testing.T) {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			[]int{2, 1},
			1,
		},
		{
			[]int{1, 2, 1, 1, 1},
			3,
		},
		{
			[]int{1},
			0,
		},
	}
	for _, testItem := range tests {
		assert.Equal(t, testItem.result, jump(testItem.nums))

	}

}
