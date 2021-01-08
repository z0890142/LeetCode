package MergeSortedArray

import (
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}{
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
		},
	}

	for _, testItem := range tests {
		merge(testItem.nums1, testItem.m, testItem.nums2, testItem.n)

	}
}
