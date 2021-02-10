package Pow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myPow(t *testing.T) {
	tests := []struct {
		x      float64
		n      int
		result float64
	}{
		{
			2.00000,
			10,
			1024.00000,
		},
		{
			2.10000,
			3,
			9.26100,
		},
		{
			2.00000,
			-2,
			0.25,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, myPow(testItem.x, testItem.n))

	}
}
