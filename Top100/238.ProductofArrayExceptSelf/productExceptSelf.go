package ProductofArrayExceptSelf

func productExceptSelf(nums []int) []int {
	l, r, result := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	l[0] = 1
	r[len(nums)-1] = 1

	for index := 1; index < len(nums); index++ {
		l[index] = nums[index-1] * l[index-1]
	}
	for index := len(nums) - 2; index >= 0; index-- {
		r[index] = nums[index+1] * r[index+1]
	}
	for index := 0; index < len(nums); index++ {
		result[index] = r[index] * l[index]
	}
	return result
}
