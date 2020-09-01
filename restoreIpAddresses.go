package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var result []string
	var tmp []string
	text := strings.Split(s, "")
	restoreIpAddressesHelper(text, tmp, &result)
	return result
}

func restoreIpAddressesHelper(s []string, preText []string, result *[]string) {
	if len(preText) > 4 {
		return
	}
	var tmp string
	for index := 0; index < len(s) && index < 3; index++ {
		if tmp == "0" {
			return
		}
		tmp += s[index]
		tmpInt, _ := strconv.Atoi(tmp)
		if tmpInt > 255 {
			return
		}
		restoreIpAddressesHelper(s[index+1:], append(preText, tmp), result)
	}
	if len(s) == 0 && len(preText) == 4 {
		*result = append(*result, strings.Join(preText, "."))
	}
}
