package main

func strStr(haystack string, needle string) int {
	if haystack == "" && needle == "" {
		return 0
	}
	result := -1
	for index := 0; index < len(haystack); index++ {
		isSuccess := check(haystack[index:], needle)
		if isSuccess {
			result = index
			break
		}
	}
	return result
}

func check(haystack string, needle string) bool {
	if len(haystack) < len(needle) {
		return false
	}
	for index, _ := range needle {
		if needle[index] != haystack[index] {
			return false
		}
	}
	return true
}
