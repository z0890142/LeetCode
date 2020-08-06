package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	combinationSumHelper(candidates, []int{}, target, &result)
	return result
}

func combinationSumHelper(candidates []int, preSlice []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, preSlice...))
		return
	}
	for i := 0; i < len(candidates) && target-candidates[i] >= 0; i++ {
		combinationSumHelper(candidates[i:], append(preSlice, candidates[i]), target-candidates[i], result)
	}
}
