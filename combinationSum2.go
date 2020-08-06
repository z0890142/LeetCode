package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	combinationSumHelper2(candidates, []int{}, target, &result)
	return result
}

func combinationSumHelper2(candidates []int, preSlice []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, preSlice...))
		return
	}

	for i := 0; i < len(candidates) && target-candidates[i] >= 0; i++ {
		combinationSumHelper2(candidates[i+1:], append(preSlice, candidates[i]), target-candidates[i], result)

		for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
			i++
		}
	}
}
