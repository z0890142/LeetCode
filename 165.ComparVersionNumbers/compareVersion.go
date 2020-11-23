package ComparVersionNumbers

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for index, val := range v1 {
		val1, _ := strconv.Atoi(val)

		if index < len(v2) {
			val2, _ := strconv.Atoi(v2[index])
			if val1 > val2 {
				return 1
			} else if val1 < val2 {
				return -1
			}
		} else {
			if val1 > 0 {
				return 1
			}
		}

	}
	for index := len(v1); index < len(v2); index++ {
		val2, _ := strconv.Atoi(v2[index])
		if val2 > 0 {
			return -1
		}
	}
	return 0
}
