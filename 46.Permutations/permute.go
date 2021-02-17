package Permutations

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
