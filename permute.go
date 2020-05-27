package main

func permuteㄍ(nums []int) [][]int {
	res := [][]int{}
	helper(nums, &res, 0)
	return res
}

func helper(nums []int, res *[][]int, i int) {
	if i == len(nums) {
		c := make([]int, len(nums))
		copy(c, nums)
		*res = append(*res, c)
	}

	for m := i; m < len(nums); m++ {
		nums[i], nums[m] = nums[m], nums[i]
		helper(nums, res, i+1)
		nums[i], nums[m] = nums[m], nums[i]
	}
}

func permute(nums []int) [][]int {
	var result [][]int
	findPermute(nums, []int{}, &result)
	return result
}

func findPermute(nums []int, tmp []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, tmp)
	}
	for index, val := range nums {
		tmpNums := append([]int{}, nums[index+1:]...)
		tmpNums = append(tmpNums, nums[0:index]...)
		findPermute(tmpNums, append(tmp, val), result)
	}
}
