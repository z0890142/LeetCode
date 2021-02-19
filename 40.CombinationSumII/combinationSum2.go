package CombinationSumII

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	combinationSum2Helper(candidates, target, &result, []int{})
	return result
}

func combinationSum2Helper(candidates []int, target int, result *[][]int, prev []int) {
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
		combinationSum2Helper(candidates[index+1:], target-candidates[index], result, append(prev, candidates[index]))
		for index+1 < len(candidates) && candidates[index] == candidates[index+1] {
			index++
		}
	}
}
