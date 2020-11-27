package Inputarrayissorted

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}
	return nil
}
