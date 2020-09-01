package main

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	var result []string
	if len(b) > len(a) {
		a, b = b, a
	}
	carry := 0
	for index := 1; index < len(a); index++ {
		if len(b)-index >= 0 {
			tmp := int(a[len(a)-index]) - 48 + int(b[len(b)-index]) - 48 + carry
			if tmp == 3 {
				result = append(result, "1")
				carry = 1
			} else if tmp == 2 {
				result = append(result, "0")
				carry = 1
			} else {
				carry = 0
				v := strconv.Itoa(tmp)
				result = append(result, v)
			}
		} else {
			tmp := int(a[len(a)-index]) - 48 + carry
			if tmp == 2 {
				result = append(result, "0")
				carry = 1
			} else {
				carry = 0
				v := strconv.Itoa(tmp)
				result = append(result, v)
			}
		}
	}
	if carry == 1 {
		result = append(result, "1")
	}
	return strings.Join(result, "")
}
