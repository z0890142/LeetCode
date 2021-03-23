package findRepeatNumber

func findRepeatNumber(nums []int) int {
	for index := 0; index < len(nums); index++ {
		for nums[index] != index {
			if nums[index] == index {
				return nums[index]
			}
			tmp := nums[index]
			nums[index], nums[tmp] = nums[tmp], nums[index]
		}
	}
	return -1
}
