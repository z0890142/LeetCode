package main

func permuteUnique(nums []int) [][]int {
	var result [][]int

	helper2(nums, &result, 0)
	return result

}

func helper2(nums []int, res *[][]int, i int) {
	if i == len(nums) {
		c := make([]int, len(nums))
		copy(c, nums)
		*res = append(*res, c)
		return
	}
	checkMap := make(map[int]int)
	for index := i; index < len(nums); index++ {
		if _, ok := checkMap[nums[index]]; ok {
			continue
		}

		nums[i], nums[index] = nums[index], nums[i]
		helper2(nums, res, i+1)
		nums[index], nums[i] = nums[i], nums[index]
		checkMap[nums[index]] = 1
	}
}
