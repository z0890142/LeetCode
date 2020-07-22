package main

import "fmt"

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	nIndex := 0
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[nIndex]
		nIndex++
	}
	if m == 0 {
		return
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1)-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			}
		}
	}
	fmt.Println(nums1)
}
