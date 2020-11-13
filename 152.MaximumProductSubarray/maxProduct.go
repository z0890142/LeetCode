package MaximumProductSubarray

func maxProduct(nums []int) int {
	minAndMax := func(a, b, c int) (int, int) {
		var min, max int
		if a < b {
			min = a
			max = b
		} else {
			min = b
			max = a
		}
		if c < min {
			min = c
		}
		if max < c {
			max = c
		}
		return min, max
	}

	var result, minDp, maxDp = nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		minDp, maxDp = minAndMax(nums[i], nums[i]*minDp, nums[i]*maxDp)
		if result < maxDp {
			result = maxDp
		}
	}
	return result
}
