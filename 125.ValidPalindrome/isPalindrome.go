package ValidPalindrome

import "strings"

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	s = strings.ToLower(s)
	for l < r {
		if !isAlnum(s[l]) {
			l++
		} else if !isAlnum(s[r]) {
			r--
		} else if s[l] != s[r] {
			return false
		} else {
			l++
			r--
		}

	}
	return true
}

func isAlnum(a byte) bool {
	return 'a' <= a && a <= 'z' || '0' <= a && a <= '9'
}
