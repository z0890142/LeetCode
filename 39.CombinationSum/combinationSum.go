package CombinationSum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	combinationSumHelper(candidates, target, &result, []int{})
	return result
}

func combinationSumHelper(candidates []int, target int, result *[][]int, prev []int) {
	if target == 0 {
		tmp := make([]int, len(prev))
		copy(tmp, prev)
		*result = append(*result, tmp)
		return
	}
	for index := 0; index < len(candidates); index++ {
		if target-candidates[index] < 0 {
			return
		}
		combinationSumHelper(candidates[index:], target-candidates[index], result, append(prev, candidates[index]))
		for index+1 < len(candidates) && candidates[index] == candidates[index+1] {
			index++
		}
	}
}
