package PalindromeNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		s      string
		result bool
	}{
		{
			"-121",
			true,
		}, {
			"112",
			false,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, isPalindrome(testItem.s))
	}

}
