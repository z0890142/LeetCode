package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		digits string
		result []string
	}{
		{
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, letterCombinations(testItem.digits))

	}

}
