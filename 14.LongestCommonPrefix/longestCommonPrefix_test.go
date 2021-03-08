package LongestCommonPrefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs   []string
		result string
	}{
		{
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			[]string{"dog", "racecar", "car"},
			"",
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, longestCommonPrefix(testItem.strs))

	}

}
