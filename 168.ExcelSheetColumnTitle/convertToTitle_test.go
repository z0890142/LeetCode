package ExcelSheetColumnTitle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convertToTitle(t *testing.T) {
	tests := []struct {
		n      int
		result string
	}{
		{
			52,

			"AZ",
		},
		{
			705,

			"AAC",
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, convertToTitle(testItem.n))
	}

}
