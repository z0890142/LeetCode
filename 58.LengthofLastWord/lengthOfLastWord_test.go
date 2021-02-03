package LengthofLastWord

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		s      string
		result int
	}{
		{

			"  ",
			0,
		},
		{
			"b   a    ",
			1,
		},
		{
			"H ",
			1,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, lengthOfLastWord(testItem.s))
	}

}
