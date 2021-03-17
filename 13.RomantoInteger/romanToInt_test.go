package RomantoInteger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		s      string
		result int
	}{
		{
			"III",
			3,
		},

		{
			"IV",
			4,
		},
		{
			"LVIII",
			58,
		},
		{
			"MCMXCIV",
			1994,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, romanToInt(testItem.s))
	}

}
