package main

func partition2(s string) [][]string {
	var result [][]string
	if len(s) == 0 {
		return result
	}
	return partition2Helper(s, []string{}, result)
}

func isPal(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true

}

func partition2Helper(s string, tmp []string, result [][]string) [][]string {
	if len(s) == 0 {
		clone := make([]string, len(tmp))
		copy(clone, tmp)
		return append(result, clone)
	}
	for index := 1; index <= len(s); index++ {
		sub_s := s[:index]
		if isPal(sub_s) {
			tmp = append(tmp, sub_s)
			result = partition2Helper(s[index:], tmp, result)
		}
	}
	return result
}
