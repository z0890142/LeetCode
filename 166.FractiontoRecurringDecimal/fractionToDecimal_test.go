package FractiontoRecurringDecimal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fractionToDecimal(t *testing.T) {
	tests := []struct {
		numerator   int
		denominator int
		result      string
	}{
		{
			4,
			333,
			"0.(012)",
		},
		{
			1,
			2,
			"0.5",
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, fractionToDecimal(testItem.numerator, testItem.denominator))
	}

}
