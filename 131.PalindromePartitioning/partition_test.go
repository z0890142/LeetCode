package PalindromePartitioning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partition(t *testing.T) {
	tests := []struct {
		s      string
		result [][]string
	}{
		{
			"aab",
			[][]string{[]string{"aa", "b"}, []string{"a", "a", "b"}},
		},
	}
	for _, testItem := range tests {
		assert.Equal(t, testItem.result, partition(testItem.s))

	}
}
