package main

import (
	"testing"
)

func Test_MergeSortedArray(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}{

		{
			[]int{0, 0, 0, 0, 0},
			0,
			[]int{-2, -1, 1, 1, 2},
			5,
		},
		{
			[]int{2, 0},
			1,
			[]int{1},
			1,
		},
		{
			[]int{0},
			0,
			[]int{1},
			1,
		},
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
		},
	}

	for _, testItem := range tests {
		MergeSortedArray(testItem.nums1, testItem.m, testItem.nums2, testItem.n)

	}

}
