package main

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for index := 0; index < len(nums); index++ {
		if index != 0 && nums[index] == nums[index-1] {
			continue
		}
		left, right := index+1, len(nums)-1

		for left < right {
			sum := nums[index] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[index], nums[left], nums[right]})
				left += 1
				right -= 1
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for right > left && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left += 1
			} else {
				right -= 1
			}
		}

	}

	return result

}
