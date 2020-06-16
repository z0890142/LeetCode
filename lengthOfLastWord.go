package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	sList := strings.Split(s, " ")
	for i := len(sList) - 1; i >= 0; i-- {
		if sList[i] != "" {
			return len(sList[i])
		}
	}
	return 0
}
