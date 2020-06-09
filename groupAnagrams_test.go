package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		str    []string
		result [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				[]string{"ate", "eat", "tea"},
				[]string{"nat", "tan"},
				[]string{"bat"},
			},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, groupAnagrams(testItem.str))
	}

}
