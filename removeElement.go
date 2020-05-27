package main

func removeElement(nums []int, val int) int {
	lIndex := 0
	for index, n := range nums {
		if n != val {
			tmp := nums[lIndex]
			nums[lIndex] = n
			nums[index] = tmp
			lIndex += 1
		}

	}
	return lIndex
}
