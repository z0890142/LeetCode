package PlusOne

func plusOne(digits []int) []int {
	var carr int

	digits[len(digits)-1] = digits[len(digits)-1] + 1
	for index := len(digits) - 1; index >= 0; index-- {
		digits[index] = digits[index] + carr
		if digits[index] > 9 {
			digits[index] = digits[index] - 10
			carr = 1
			continue
		}
		carr = 0
		break
	}
	if carr == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
