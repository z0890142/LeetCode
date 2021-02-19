package CountandSay

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	result := "1"
	for index := 2; index <= n; index++ {
		count := 1
		tmp := ""
		for index := 1; index < len(result); index++ {
			if result[index] != result[index-1] {
				tmp += strconv.Itoa(count) + result[index-1:index]
				count = 1
				continue
			}
			count++
		}
		tmp += strconv.Itoa(count) + result[len(result)-1:]
		result = tmp
	}
	return result
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	str := countAndSay(n - 1)

	count := 0
	val := "default"
	strBuilder := new(strings.Builder)

	length := len(str)

	for i := 0; i < length; i++ {
		c := string(str[i])
		if val != c {
			lastVal := val
			lastCount := count
			val = c
			count = 1

			if lastVal != "default" {
				strBuilder.WriteString(fmt.Sprintf("%d%s", lastCount, lastVal))
			}
		} else {
			count += 1
		}

		if i == length-1 {
			strBuilder.WriteString(fmt.Sprintf("%d%s", count, val))
		}
	}
	return strBuilder.String()

}
