package ReverseWordsinaString

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		s      string
		result string
	}{
		// {
		// 	"a good   example",
		// 	"example good a",
		// },

		{
			"  hello world  ",
			"world hello",
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, reverseWords(testItem.s))
	}

}
