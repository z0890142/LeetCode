package main

import "strconv"

func grayCode(n int) []int {
	var result []int
	grayCodeHelper("", n, &result)
	return result
}

func grayCodeHelper(text string, n int, result *[]int) {
	if len(text) == n {
		tmp, _ := strconv.ParseInt(text, 2, 64)
		*result = append(*result, int(tmp))
		return
	}
	grayCodeHelper(text+"0", n, result)
	grayCodeHelper(text+"1", n, result)

}
