package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		x      int
		result int
	}{
		{
			16,
			2,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, mySqrt(testItem.x))

	}
}
