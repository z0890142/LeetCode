package DivideTwoIntegers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divide(t *testing.T) {
	tests := []struct {
		dividend int
		divisor  int
		result   int
	}{
		// {
		// 	10,
		// 	3,
		// 	3,
		// },
		{
			1,
			1,
			1,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, divide(testItem.dividend, testItem.divisor))
	}

}
