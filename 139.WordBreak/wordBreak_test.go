package WordBreak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		result   bool
	}{
		{
			"aaab",
			[]string{"a", "aa"},
			false,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, wordBreak(testItem.s, testItem.wordDict))
	}

}
