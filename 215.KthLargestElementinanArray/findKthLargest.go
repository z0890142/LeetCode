package KthLargestElementinanArray

// func findKthLargest(nums []int, k int) int {
// 	l := 0
// 	r := len(nums)

// 	index := r - k
// 	for {
// 		p := partition(nums, l, r)
// 		if index == p {
// 			return nums[p]
// 		} else if index < p {
// 			r = p
// 		} else {
// 			l = p + 1
// 		}
// 	}
// }

// func partition(nums []int, l int, r int) int {

// 	v := nums[l]
// 	less := l

// 	for i := l + 1; i < r; i++ {
// 		if nums[i] < v {
// 			less++
// 			nums[less], nums[i] = nums[i], nums[less]
// 		}
// 	}
// 	nums[less], nums[l] = nums[l], nums[less]
// 	return less
// }

func findKthLargest(nums []int, k int) int { // [3,1,4,2]
	quickSort(nums, 0, len(nums)-1, k) // [3,1,4,2], 0, 3
	return nums[k-1]
}

func quickSort(nums []int, left, right int, k int) {
	if left < right {
		partitionIdx := partition(nums, left, right) // [3,2,4,1]
		// 2
		if partitionIdx > k-1 {
			quickSort(nums, left, partitionIdx, k) // [3,2,4]
		} else {
			quickSort(nums, left, partitionIdx, k)    // [3,2,4]
			quickSort(nums, partitionIdx+1, right, k) // [3,2,4]
		}
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[(left+right)/2]
	for {
		for nums[left] > pivot {
			left++
		}
		for nums[right] < pivot {
			right--
		}
		if left >= right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return right
}
