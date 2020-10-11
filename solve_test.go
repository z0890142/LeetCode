package main

import (
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		board [][]byte
	}{
		{
			[][]byte{
				[]byte{'X', 'X', 'X', 'X'},
				[]byte{'X', 'O', 'O', 'X'},
				[]byte{'X', 'X', 'O', 'X'},
				[]byte{'X', 'O', 'X', 'X'},
			},
		},
	}

	for _, testItem := range tests {
		solve(testItem.board)
	}

}
