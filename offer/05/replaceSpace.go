package replaceSpace

func replaceSpace(s string) string {
	result := ""
	for index := 0; index < len(s); index++ {
		if s[index] == ' ' {
			result += "%20"
		} else {
			result += s[index : index+1]
		}
	}
	return result
}
