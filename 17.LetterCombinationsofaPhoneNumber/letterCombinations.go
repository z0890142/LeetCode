package LetterCombinationsofaPhoneNumber

var numMap = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var result []string
	for index := 0; index < len(digits); index++ {
		if len(result) == 0 {
			result = append(result, numMap[digits[index:index+1]]...)
			continue
		}
		var tmp []string
		for _, s1 := range result {
			for _, s2 := range numMap[digits[index:index+1]] {
				tmp = append(tmp, s1+s2)
			}
		}
		result = tmp
	}
	return result
}

// func letterCombinations(digits string) []string {
// 	var result []string
//     if digits == "" {
// 		return result
// 	}
// 	findCombinations(digits, "", &result)
// 	return result
// }

// func findCombinations(digits string, preText string, result *[]string) {
// 	if digits == "" {
// 		*result = append(*result, preText)
// 		return
// 	}
// 	for _, char := range numMap[digits[0:1]] {
// 		findCombinations(digits[1:], preText+char, result)
// 	}
// }
