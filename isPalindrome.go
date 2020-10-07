package main

var diff byte = 'a' - 'A'

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if !isAlnum(s[l]) {
			l++
		} else if !isAlnum(s[r]) {
			r--
		} else if toLower(s[l]) != toLower(s[r]) {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}

func isAlnum(a byte) bool {
	return 'a' <= toLower(a) && toLower(a) <= 'z' || '0' <= a && a <= '9'
}

func toLower(a byte) byte {
	if a >= 'A' && a <= 'Z' {
		return a + diff
	}
	return a
}
