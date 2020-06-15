package main

import "sort"

func merge(intervals [][]int) [][]int {
	var result [][]int
	if len(intervals) <= 1 {
		return intervals
	}
	var start []int
	var end []int
	for _, it := range intervals {
		start = append(start, it[0])
		end = append(end, it[1])
	}
	sort.Ints(start)
	sort.Ints(end)

	for startIndex, endIndex := 0, 0; endIndex < len(end); endIndex++ {
		if endIndex == len(end)-1 || start[endIndex+1] > end[endIndex] { //中间出现了不连续，记录前一个连续值
			result = append(result, []int{start[startIndex], end[endIndex]})
			startIndex = endIndex + 1
		}
	}

	return result
}
