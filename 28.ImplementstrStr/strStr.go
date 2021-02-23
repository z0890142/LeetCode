package ImplementstrStr

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	i, j, k := 0, 0, 0

	for i < len(haystack) {
		if j >= len(needle) {
			return k
		}

		if haystack[i] == needle[j] {
			i++
			j++
		} else if haystack[i] != needle[j] {
			k += 1
			i = k
			j = 0
		}
	}

	if i >= len(haystack) && j >= len(needle) {
		return k
	}
	return -1
}
