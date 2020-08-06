package main

import (
	"strconv"
)

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	var result int
	numDecodingsHelper(s, &result)
	return result
}

func numDecodingsHelper(s string, result *int) {

	for index := 1; index <= len(s); index++ {
		intVal, _ := strconv.Atoi(s[0:index])
		if intVal <= 0 || intVal > 26 {
			break
		}
		if index-1 == len(s)-1 {
			*result += 1
			break
		}
		numDecodingsHelper(s[index:], result)
	}
}
