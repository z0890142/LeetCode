package NumberofIslands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numIslands(t *testing.T) {
	tests := []struct {
		grid   [][]byte
		result int
	}{
		{
			[][]byte{
				[]byte{'1', '1', '0', '0', '0'},
				[]byte{'1', '1', '0', '0', '0'},
				[]byte{'0', '0', '1', '0', '0'},
				[]byte{'0', '0', '0', '1', '1'},
			},
			3,
		},
		{
			[][]byte{
				[]byte{'1', '1', '1', '1', '0'},
				[]byte{'1', '1', '0', '1', '0'},
				[]byte{'1', '1', '0', '0', '0'},
				[]byte{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			[][]byte{
				[]byte{'1', '1', '1', '1', '0'},
				[]byte{'1', '1', '0', '1', '0'},
				[]byte{'1', '1', '0', '0', '0'},
				[]byte{'0', '0', '0', '0', '1'},
			},
			2,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, numIslands(testItem.grid))
	}

}
