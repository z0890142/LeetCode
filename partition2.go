package main

func partition2(s string) [][]string {

	var result [][]string

	for startIndex := 0; startIndex < len(s); startIndex++ {
		var tmp []string
		isPalindrome, endIndex := partition2Helper(startIndex, s)
		if !isPalindrome {
			tmp = append(tmp, s[startIndex:startIndex+1])
			continue
		}
		tmp = append(tmp, s[startIndex:endIndex+1])
		endIndex++
		for endIndex < len(s) {
			tmp = append(tmp, s[endIndex:endIndex+1])
			endIndex++
		}
		result = append(result, tmp)
	}
	return result
}

func partition2Helper(startIndex int, s string) (isPalindrome bool, end int) {

	for endPoint := len(s) - 1; endPoint > startIndex; endPoint-- {
		isPalindrome = true
		for endIndex := endPoint; endPoint > startIndex; endIndex-- {
			if s[startIndex] != s[endIndex] {
				isPalindrome = false
				break
			}
			startIndex++
		}
		if isPalindrome {
			return isPalindrome, endPoint
		}
	}

	return false, 0
}
