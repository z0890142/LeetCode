package main

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return countNum(2, "1", n)
}

func countNum(num int, text string, target int) string {
	if num > target {
		return text
	}

	count := 1
	result := ""
	for index := len(text) - 1; index > 0; index-- {
		if text[index] != text[index-1] {
			result = strconv.Itoa(count) + text[index:index+1] + result
			count = 1
		} else {
			count++
		}
	}

	result = strconv.Itoa(count) + text[0:1] + result

	return countNum(num+1, result, target)
}
