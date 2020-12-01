package PalindromePartitioning

func partition(s string) [][]string {
	var result [][]string
	if len(s) == 0 {
		return result
	}
	return partitionHelper(s, []string{}, result)

}

func isPal(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true

}

func partitionHelper(s string, tmp []string, result [][]string) [][]string {
	if len(s) == 0 {
		clone := make([]string, len(tmp))
		copy(clone, tmp)
		return append(result, clone)
	}
	for index := 1; index <= len(s); index++ {
		sub_s := s[:index]
		if isPal(sub_s) {
			tmp = append(tmp, sub_s)
			result = partitionHelper(s[index:], tmp, result)
		}
	}
	return result
}
