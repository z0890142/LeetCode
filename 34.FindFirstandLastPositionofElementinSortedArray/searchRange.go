package FindFirstandLastPositionofElementinSortedArray

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	var mid int

	for {
		if l > r {
			return []int{-1, -1}
		}

		mid = (l + r) / 2

		if nums[mid] == target {
			break
		}

		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	l, r = mid, mid
	for ; l >= 0 && nums[l] == target; l-- {
	}
	for ; r < len(nums) && nums[r] == target; r++ {
	}

	return []int{l + 1, r - 1}
}
