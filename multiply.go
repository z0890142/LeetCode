package main

import "strconv"

func multiply(num1 string, num2 string) string {
	var resultStr string
	result := make([]int, len(num1)+len(num2))
	if len(num2) > len(num1) {
		num1, num2 = num2, num1
	}
	for num2Index := len(num2) - 1; num2Index >= 0; num2Index-- {
		for num1Index := len(num1) - 1; num1Index >= 0; num1Index-- {
			num2Val := int(num2[num2Index]) - '0'
			num1Val := int(num1[num1Index]) - '0'
			sum := num2Val * num1Val
			result[num2Index+num1Index+1] += sum
			result[num2Index+num1Index] += result[num2Index+num1Index+1] / 10
			result[num2Index+num1Index+1] = result[num2Index+num1Index+1] % 10
		}
	}
	start := false
	for i := 0; i < len(result); i++ {
		if result[i] != 0 {
			start = true
		}
		if start {
			resultStr += strconv.Itoa(result[i])
		}
	}
	if resultStr == "" {
		return "0"
	}
	return resultStr
}
