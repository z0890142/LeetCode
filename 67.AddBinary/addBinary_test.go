package AddBinary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addBinary(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		result string
	}{
		{
			"100",
			"110010",
			"110110",
		},
		{
			"11",
			"1",
			"100",
		},
		{
			"1010",
			"1011",
			"10101",
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, addBinary(testItem.a, testItem.b))

	}

}
