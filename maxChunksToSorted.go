package main

func maxChunksToSorted(arr []int) int {
	var result int
	cutIndex := -1
	for index := 0; index < len(arr); index++ {
		if arr[index] == index && cutIndex == -1 {
			result++
		} else {
			if arr[index] > cutIndex {
				cutIndex = arr[index]
				continue
			}
			if index == cutIndex {
				result++
				cutIndex = -1
			}
		}
	}
	return result
}
