package ClimbingStairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		n      int
		result int
	}{
		{
			2,
			2,
		},
		{
			3,
			3,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, climbStairs(testItem.n))

	}

}
