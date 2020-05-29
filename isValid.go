package main

func isValid(s string) bool {
	vaildMap := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	var checkList []rune
	for _, char := range s {
		if val, ok := vaildMap[char]; ok {
			checkList = append([]rune{val}, checkList...)
			continue
		}
		if len(checkList) == 0 {
			return false
		}
		if checkList[0] != char {
			return false
		}
		checkList = checkList[1:]
	}
	if len(checkList) == 0 {
		return true

	}
	return false

}
