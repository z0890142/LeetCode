package main

import "fmt"

func Solution(S string) int {
	// write your code in Go 1.4
	var val int
	var multiple = 1
	var result int
	var repeatTime int

	for index := len(S) - 1; index >= 0; index-- {
		if S[index] == '1' {
			repeatTime++
		}
		val = val + int(S[index]-'0')*multiple
		multiple *= 2
	}

	if repeatTime > 400000 {
		return 799999
	}

	for val > 0 {
		result++
		if val%2 == 0 {
			val = val / 2
		} else {
			val -= 1
		}

	}
	return result
}

func main() {
	fmt.Println(Solution("1111010101111"))
}
