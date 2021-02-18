package MultiplyStrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_multiply(t *testing.T) {

	tests := []struct {
		num1   string
		num2   string
		result string
	}{
		// {
		// 	"0",
		// 	"0",
		// 	"0",
		// },
		{
			"123",
			"456",
			"56088",
		},
	}

	for _, testItem := range tests {

		assert.Equal(t, testItem.result, multiply(testItem.num1, testItem.num2))

	}
}
