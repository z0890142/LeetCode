package MajorityElement

import "sort"

func majorityElement(nums []int) int {
	checkMap := make(map[int]int)
	l := len(nums) / 2
	sort.Ints(nums)
	for _, val := range nums {
		if _, isExist := checkMap[val]; isExist {
			checkMap[val] += 1
		} else {
			checkMap[val] = 1
		}
		if checkMap[val] > l {
			return val
		}
	}
	return -1
}
