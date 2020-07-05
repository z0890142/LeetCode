package main

func plusOne(digits []int) []int {
	digits[len(digits)-1] += 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]/10 > 0 && i != 0 {
			digits[i-1] += digits[i] / 10
			digits[i] = digits[i] % 10
		} else if digits[i]/10 > 0 {
			tmp := []int{digits[i] / 10}
			digits[i] = digits[i] % 10
			digits = append(tmp, digits...)
		}
	}
	return digits
}
