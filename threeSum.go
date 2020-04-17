package main

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	// for index1 := 0; index1 < len(nums); index1++ {
	// 	for index2 := index1 + 1; index2 < len(nums); index2++ {
	// 		if nums[index1] > nums[index2] {
	// 			tmp := nums[index1]
	// 			nums[index1] = nums[index2]
	// 			nums[index2] = tmp
	// 		}
	// 	}
	// }
	sort.Ints(nums)

	for index, val := range nums {
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		start := index + 1
		end := len(nums) - 1
		for start < end {
			sum := val + nums[start] + nums[end]
			if sum == 0 {
				result = append(result, []int{val, nums[start], nums[end]})
				start++
				end--
				for start < end && nums[start] == nums[start-1] {
					start++
				}
				for start < end && nums[end] == nums[end+1] {
					end--
				}
			}
			if sum < 0 {
				start++
			} else if sum > 0 {
				end--
			}
		}
	}
	return result
}
