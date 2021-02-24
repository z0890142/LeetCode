package RemoveElement

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[r] == val {
			r--
			continue
		}
		if nums[l] == val {
			nums[l], nums[r] = nums[r], nums[l]
			r--
			l++
			continue
		}
		l++
	}
	return l
}

// func removeElement(nums []int, val int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	index := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] != val {
// 			nums[index] = nums[i]
// 			index++
// 		}

// 	}
// 	return index
// }
