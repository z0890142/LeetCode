package FindPeakElement

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] > nums[1] {
		return 0
	}

	l := 0
	r := len(nums) - 1
	if nums[r] > nums[r-1] {
		return r
	}
	mid := (l + r) / 2
	for mid >= 1 && nums[mid] < nums[mid-1] {
		mid--
	}
	if mid != (l+r)/2 {
		return mid
	}
	for mid <= r-1 && nums[mid] < nums[mid+1] {
		mid++
	}

	return mid
}

func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
