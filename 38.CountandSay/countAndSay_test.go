package CountandSay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countAndSay(t *testing.T) {
	tests := []struct {
		n      int
		result string
	}{
		{
			1,
			"1",
		},
		{
			4,
			"1211",
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, countAndSay(testItem.n))

	}

}
