package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for start := 0; start < len(nums)-3; start++ {
		if start != 0 && nums[start-1] == nums[start] {
			continue
		}
		for index := start + 1; index < len(nums); index++ {
			if index != start+1 && nums[index-1] == nums[index] {
				continue
			}
			left, right := index+1, len(nums)-1

			for left < right {
				sum := nums[start] + nums[index] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[start], nums[index], nums[left], nums[right]})
					left += 1
					right -= 1
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for right > left && nums[right] == nums[right+1] {
						right--
					}

				} else if sum < target {
					left += 1
				} else {
					right -= 1
				}
			}

		}
	}

	return result
}
