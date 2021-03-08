package LongestCommonPrefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var result []byte
	result = []byte(strs[0])
	for index := 1; index < len(strs); index++ {
		if len(result) == 0 {
			return ""
		}

		tmp := strs[index]
		if len(tmp) > len(result) {
			tmp = tmp[:len(result)]
		} else if len(tmp) < len(result) {
			result = result[:len(tmp)]
		}

		for i := len(result) - 1; i >= 0; i-- {
			if tmp[i] != result[i] {
				result = result[:i]
			}
		}

	}
	return string(result)
}

// func longestCommonPrefix(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}
// 	result := strs[0]
// 	charIndex := 0
// 	isStop := false
// 	for !isStop {
// 		for _, str := range strs {
// 			if charIndex == len(str) {
// 				isStop = true
// 				break
// 			}
// 			if result[charIndex] != str[charIndex] {
// 				isStop = true
// 				break
// 			}
// 		}
// 		charIndex += 1
// 	}
// 	result = result[:charIndex-1]
// 	return result
// }
