package fourSum

import (
	"sort"
)

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

// func fourSum(nums []int, target int) [][]int {
//     sort.Ints(nums)
//     return kSum(nums, 0, 4, target)
// }

// func kSum(nums []int, start, k, target int) (r [][]int) {
//     if start == len(nums) || nums[start] * k > target || target > nums[len(nums)-1]*k {
//         return
//     }
//     if k == 2 {
//         return twoSum(nums, start, target)
//     }
//     for i := start; i < len(nums); i++ {
//         if i != start && nums[i] == nums[i-1] {
//             continue
//         }
//         t := kSum(nums, i+1, k-1, target - nums[i])
//         for _, group := range t {
//             r = append(r, append(group, nums[i]))
//         }
//     }
//     return
// }

// func twoSum(nums []int, start, target int) (r [][]int) {
//     fmt.Println("in twoSum", start, target)
//     lo, hi := start, len(nums)-1
//     for lo < hi {
//         sum := nums[lo] + nums[hi]
//         if sum < target {
//             lo++
//             continue
//         } else if sum > target || (lo > start && nums[lo] == nums[lo-1]) {
//             hi--
//             continue
//         }
//         r = append(r, []int{nums[lo], nums[hi]})
//         lo++
//         hi--
//     }
//     return
// }
