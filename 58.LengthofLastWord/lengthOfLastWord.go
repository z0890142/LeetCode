package LengthofLastWord

func lengthOfLastWord(s string) int {
	var res int
	if len(s) > 0 {
		var end int = len(s) - 1
		for end >= 0 && s[end] == ' ' {
			end--
		}
		for i := end; i >= 0; i-- {
			if s[i] != ' ' {
				res++
			} else {
				break
			}
		}
	}
	return res
}
