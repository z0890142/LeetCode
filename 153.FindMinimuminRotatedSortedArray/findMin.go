package FindMinimuminRotatedSortedArray

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	mid := len(nums) / 2
	m1, m2 := findMin(nums[:mid]), findMin(nums[mid:])
	if m1 > m2 {
		return m2
	}
	return m1
}
