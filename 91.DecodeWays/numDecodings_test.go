package DecodeWays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numDecodings(t *testing.T) {
	tests := []struct {
		s      string
		result int
	}{
		{
			"12",
			2,
		},
		{
			"226",
			3,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, numDecodings(testItem.s))

	}
}
