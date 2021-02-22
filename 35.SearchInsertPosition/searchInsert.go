package SearchInsertPosition

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	for i := range nums {
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			return i
		}
	}

	return len(nums)
}
