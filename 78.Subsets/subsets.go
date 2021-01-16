package Subsets

func subsets(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})
	subsetsHelper(nums, []int{}, &result)
	return result
}

func subsetsHelper(nums []int, preSlices []int, result *[][]int) {
	for index, val := range nums {

		preSlices = append(preSlices, val)
		tmp := make([]int, len(preSlices))
		copy(tmp, preSlices)
		*result = append(*result, tmp)

		subsetsHelper(nums[index+1:], preSlices, result)
		preSlices = preSlices[0 : len(preSlices)-1]
	}

}
