package RestoreIPAddresses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_restoreIpAddresses(t *testing.T) {
	tests := []struct {
		s      string
		result []string
	}{
		{
			"25525511135",
			[]string{
				"255.255.11.135",
				"255.255.111.35",
			},
		},
		{
			"0000",
			[]string{
				"0.0.0.0",
			},
		}, {
			"1111",
			[]string{
				"1.1.1.1",
			},
		}, {
			"010010",
			[]string{
				"0.10.0.10",
				"0.100.1.0",
			},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, restoreIpAddresses(testItem.s))

	}
}
