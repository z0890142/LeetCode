package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		s      string
		result bool
	}{
		{
			"(){}[]",
			true,
		}, {
			"{()}[]",
			true,
		}, {
			"{()}[]",
			true,
		}, {
			"{]",
			false,
		}, {
			"[{]}",
			false,
		}, {
			"[",
			false,
		}, {
			"",
			true,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, isValid(testItem.s))

	}
}
