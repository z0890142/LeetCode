package ReverseWordsinaString

import "strings"

func reverseWords(s string) string {
	text_slice := strings.Split(s, " ")
	result := []string{}
	for index := len(text_slice) - 1; index >= 0; index-- {
		if text_slice[index] == "" {
			continue
		}
		result = append(result, text_slice[index])
	}

	return strings.Join(result, " ")
}
