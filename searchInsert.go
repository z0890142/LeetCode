package main

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1

	if nums[0] >= target {
		return 0
	} else if nums[len(nums)-1] == target {
		return len(nums) - 1

	} else if nums[len(nums)-1] < target {
		return len(nums)
	}

	for l <= r {
		mid := (r + l) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
