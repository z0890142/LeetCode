package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		n      int
		result []string
	}{
		{
			3,
			[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, generateParenthesis(testItem.n))

	}

}
