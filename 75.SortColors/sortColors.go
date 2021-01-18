package SortColors

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	length := len(nums)

	redIndex, blueIndex := 0, length-1
	for index := 0; index < blueIndex; {
		if nums[index] == 2 {
			nums[index], nums[blueIndex] = nums[blueIndex], nums[index]
			blueIndex--
		} else if nums[index] == 0 {
			nums[index], nums[redIndex] = nums[redIndex], nums[index]
			redIndex++
			index++
		} else {
			index++
		}
	}
}
