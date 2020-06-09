package main

import (
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		matrix [][]int
		result [][]int
	}{
		{
			[][]int{[]int{5, 1, 9, 11}, []int{2, 4, 8, 10}, []int{13, 3, 6, 7}, []int{15, 14, 12, 16}},
			[][]int{[]int{15, 13, 2, 5}, []int{14, 3, 4, 1}, []int{12, 6, 8, 9}, []int{16, 7, 10, 11}},
		},
	}

	for _, testItem := range tests {
		rotate(testItem.matrix)

	}
}
