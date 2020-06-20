package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getPermutation(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		result string
	}{
		{
			3,
			3,
			"213",
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, getPermutation(testItem.n, testItem.k))

	}
}
