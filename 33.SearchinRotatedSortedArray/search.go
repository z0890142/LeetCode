package SearchinRotatedSortedArray

// func search(nums []int, target int) int {
// 	if nums[0] == target {
// 		return 0
// 	}
// 	if nums[len(nums)-1] > nums[0] {
// 		return find(nums, target, 0, len(nums)-1)
// 	}
// 	for index := 1; index < len(nums); index++ {
// 		if nums[index] < nums[index-1] {
// 			if nums[0] <= target {
// 				return find(nums, target, 0, index-1)
// 			} else if nums[index] <= target {
// 				return find(nums, target, index, len(nums)-1)
// 			}
// 		}
// 	}
// 	return -1
// }

// func find(nums []int, target int, l int, r int) int {
// 	var mid int

// 	for {
// 		if l > r {
// 			return -1
// 		}

// 		mid = (l + r) / 2

// 		if nums[mid] == target {
// 			break
// 		}

// 		if nums[mid] > target {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return mid
// }

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if (nums[r] < nums[mid] && (target > nums[mid] || target <= nums[r])) || (nums[r] >= target && target > nums[mid]) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
