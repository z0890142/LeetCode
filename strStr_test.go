package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_strStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		result   int
	}{
		{
			"hello",
			"ll",
			2,
		},
		{
			"aaaaaa",
			"bba",
			-1,
		},
		{
			"test",
			"s",
			2,
		},
		{
			"",
			"",
			0,
		},
		{
			"aa",
			"aaa",
			-1,
		},
		{
			"mississippi",
			"issipi",
			-1,
		},
	}
	for _, testItem := range tests {
		assert.Equal(t, testItem.result, strStr(testItem.haystack, testItem.needle))

	}
}
