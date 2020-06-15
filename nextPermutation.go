package main

func nextPermutation(nums []int) {
	targetIndex := 0
	for index := len(nums) - 1; index > 0; index-- {
		if nums[index-1] >= nums[index] {
			continue
		}
		targetIndex = index - 1
		for index2 := len(nums) - 1; index2 > targetIndex; index2-- {
			if nums[index2] > nums[targetIndex] {
				swap(&nums[targetIndex], &nums[index2])
				break
			}
		}
		sortNums(nums, targetIndex+1, len(nums)-1)
		return
	}
	sortNums(nums, 0, len(nums)-1)
}

func swap(num1 *int, num2 *int) {
	tmp := *num1
	*num1 = *num2
	*num2 = tmp
}

func sortNums(nums []int, l int, r int) {

	for l < r {
		swap(&nums[l], &nums[r])
		l += 1
		r -= 1
	}
}
