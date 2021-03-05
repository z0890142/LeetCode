package threeSum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for index := 0; index < len(nums)-2; index++ {
		if index != 0 && nums[index-1] == nums[index] {
			continue
		}
		l, r := index+1, len(nums)-1
		for l < r {
			sum := nums[index] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[index], nums[l], nums[r]})
				l += 1
				r -= 1
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for r > l && nums[r] == nums[r+1] {
					r--
				}
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return result
}
