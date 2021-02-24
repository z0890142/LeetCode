package RemoveDuplicatesfromSortedArray

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 1
	target := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != target {
			nums[index] = nums[i]
			target = nums[i]
			index++
		}
	}
	return index
}

// func removeDuplicates(nums []int) int {
// 	x := 1
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i - 1] != nums[i] {
// 			nums[x] = nums[i]
// 			x++
// 		}
// 	}
// 	if len(nums) != 0 {
// 		nums = nums[:x]
// 		return x
// 	} else {
// 		return 0
// 	}
// }
