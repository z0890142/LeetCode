package ValidPalindrome

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
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"race a car",
			false,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, isPalindrome(testItem.s))
	}

}
