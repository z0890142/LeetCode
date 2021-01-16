package WordSearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_test(t *testing.T) {
	tests := []struct {
		board  [][]byte
		word   string
		result bool
	}{
		{
			[][]byte{[]byte{'C', 'A', 'A'}, []byte{'A', 'A', 'A'}, []byte{'B', 'C', 'D'}},
			"AAB",
			true,
		},
		{
			[][]byte{
				[]byte{'A', 'B', 'C', 'E'},
				[]byte{'S', 'F', 'E', 'S'},
				[]byte{'A', 'D', 'E', 'E'}},
			"ABCEFSADEESE",
			true,
		},
		{
			[][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}},
			"ABCCED",
			true,
		},
		{
			[][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}},
			"SEE",
			true,
		},
		{
			[][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}},
			"ABCB",
			false,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, exist(testItem.board, testItem.word))

	}

}
