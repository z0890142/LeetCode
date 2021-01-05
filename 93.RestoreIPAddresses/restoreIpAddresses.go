package RestoreIPAddresses

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var result []string
	restoreIpAddressesHelper(s, []string{}, &result)
	return result
}
func restoreIpAddressesHelper(s string, sub []string, result *[]string) {
	if len(sub) == 4 && len(s) > 0 {
		return
	}
	if len(sub) == 4 {

		*result = append(*result, strings.Join(sub, "."))
		return
	}
	for index := 1; index <= len(s) && index <= 3; index++ {
		tmpString := s[0:index]

		if index > 1 && tmpString[0] == '0' {
			return
		}
		tmpInt, _ := strconv.Atoi(tmpString)
		if tmpInt > 255 {
			return
		}
		restoreIpAddressesHelper(s[index:], append(sub, tmpString), result)
	}
}
