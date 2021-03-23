package ProductofArrayExceptSelf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, productExceptSelf(testItem.nums))
	}

}
